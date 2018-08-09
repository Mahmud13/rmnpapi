package controllers

import (
	"github.com/astaxie/beego"
  "log"

  "golang.org/x/net/context"

  firebase "firebase.google.com/go"

  "firebase.google.com/go/messaging"

  "google.golang.org/api/option"
)
type TestController struct {
  beego.Controller
}
// @Title GetAll
// @Description get all objects
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router / [get]
func (c *TestController) Index() {
  opt := option.WithCredentialsFile("storage/robi-mnp-firebase-adminsdk-eyz6a-6a47d7b2bb.json")
  ctx := context.Background()
  app, err := firebase.NewApp(ctx, nil, opt)
  if err != nil {
    log.Fatalf("error initializing app: %v\n", err)
  }
  // Access auth service from the default app
  client, err := app.Messaging(ctx)
  if err != nil {
    log.Fatalf("error getting Auth client: %v\n", err)
  }
  // This registration token comes from the client FCM SDKs.
  registrationToken := "d1Qe0B0hxaM:APA91bGUHmoaUyHxzRVbbkNRn7KDgrT73yEZvgiGvFKRyqvjvlv8D0kSj17NrDJwEb8xdqfDcki4cWW_vjjiuGHODsMgXJrY_qDzdj9ZSYB_N01r24fGgtEkleFpCn5I3hi1jZUBclVuaDPI-6S4aG4TRR-PzGdvRA"

  // See documentation on defining a message payload.
  message := &messaging.Message{
        Data: map[string]string{
                "title": "850",
                "body":  "2:45",
        },
        Token: registrationToken,
    }
    // Send a message to the device corresponding to the provided
    // registration token.
    response, err := client.Send(ctx, message)
    if err != nil {
      c.Data["json"] = err
    }else{
      c.Data["json"] = response
    }
    c.ServeJSON()
}
