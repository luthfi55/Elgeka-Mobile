package initializers

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mdp/qrterminal/v3"
	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types"
	waLog "go.mau.fi/whatsmeow/util/log"
	"google.golang.org/protobuf/proto"
)

var client *whatsmeow.Client
var clientMutex sync.Mutex

func SendMessageToUser(userNumber string, otp string) {
	phoneNumber := userNumber
	server := "s.whatsapp.net"
	senderJID := types.NewJID(phoneNumber, server)
	message := &waProto.Message{
		Conversation: proto.String("Kode OTP Anda: " + otp + ". Jangan berikan kode ini kepada orang lain."),
	}
	client.SendMessage(context.Background(), senderJID, message)
}

func ConnectToWhatsapp() error {
	dbLog := waLog.Stdout("Database", "DEBUG", true)
	container, err := sqlstore.New("sqlite3", "file:whatsapp.db?_foreign_keys=on", dbLog)
	if err != nil {
		return fmt.Errorf("failed to create sqlstore: %v", err)
	}

	deviceStore, err := container.GetFirstDevice()
	if err != nil {
		return fmt.Errorf("failed to get first device: %v", err)
	}

	clientLog := waLog.Stdout("Client", "DEBUG", true)

	clientMutex.Lock()
	client = whatsmeow.NewClient(deviceStore, clientLog)
	clientMutex.Unlock()

	if client.Store.ID == nil {
		qrChan, _ := client.GetQRChannel(context.Background())
		err = client.Connect()
		if err != nil {
			return fmt.Errorf("failed to connect: %v", err)
		}
		for evt := range qrChan {
			if evt.Event == "code" {
				qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)
			} else {
				fmt.Println("Login event:", evt.Event)
			}
		}
	} else {
		err = client.Connect()
		if err != nil {
			return fmt.Errorf("failed to connect: %v", err)
		}
	}

	log.Println("Connected to WhatsApp")
	return nil
}

func DisconnectWhatsapp() {
	clientMutex.Lock()
	defer clientMutex.Unlock()

	if client != nil {
		client.Disconnect()
		log.Println("Disconnected from WhatsApp")
	}
}
