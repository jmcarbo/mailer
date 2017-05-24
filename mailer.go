package main

import (
  "os"
  "github.com/codegangsta/cli"
  "github.com/jordan-wright/email"
  "log"
  "path"
  "net/smtp"
)

func sendMail(user, password, server, port, from, to, subject, message, fileName string) {
  if to == "" {
    log.Println("No destination")
    return
  }
  e := email.NewEmail()
  e.From = from
  e.To = []string{to}
  //e.Bcc = []string{"test_bcc@example.com"}
  //e.Cc = []string{"test_cc@example.com"}
  e.Subject = subject
  e.Text = []byte(message)
  //e.HTML = []byte("<h1>Fancy HTML is supported, too!</h1>")
  if fileName != "" {
    filepath, _ := os.Getwd()
    _, err := e.AttachFile(path.Join(filepath, fileName))
    if err != nil {
      log.Println(err)
    } else {
      log.Println("Attaching file " + path.Join(filepath, fileName))
    }
  }

  err := e.Send(server + ":" + port, smtp.PlainAuth("", user, password, server))
  if err != nil {
    log.Println(err)
  } else {
    log.Println("All good brah")
  }
}

func main() {
  app := cli.NewApp()
  app.Name = "mailer"
  app.Usage = "mailer -user ..."
  app.Action = func(c *cli.Context) {
    println("Sending Mail")
    sendMail(c.String("user"), c.String("password"), c.String("server"), 
      c.String("port"), c.String("from"), c.String("to"), c.String("subject"), c.String("message"), c.String("filename"))
  }

  app.Flags = []cli.Flag {
    cli.StringFlag{
      Name: "user",
      Value: "user@gmail.com",
      Usage: "Server username",
    },
    cli.StringFlag{
      Name: "password",
      Value: "a password",
      Usage: "..",
    },
    cli.StringFlag{
      Name: "server",
      Value: "smtp.gmail.com",
      Usage: "..",
    },
    cli.StringFlag{
      Name: "port",
      Value: "587",
      Usage: "..",
    },
    cli.StringFlag{
      Name: "from",
      Value: "user@gmail.com",
      Usage: "<from@emailc.com>",
    },
    cli.StringFlag{
      Name: "to",
      Value: "",
      Usage: "..",
    },
    cli.StringFlag{
      Name: "subject",
      Value: "No subject",
      Usage: "<subject>",
    },
    cli.StringFlag{
      Name: "message",
      Value: "No message",
      Usage: "..",
    },
    cli.StringFlag{
      Name: "filename",
      Value: "",
      Usage: "..",
    },
  }
  app.Run(os.Args)
}
