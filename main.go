package main

import (
  "fmt"
  "../reddit"
  "time"
  "strings"
)

func main() {
  r, err := reddit.NewLoginSession("FuckTheCowboysBot", "im1coolguy", "FuckTheCowboysBot/v0.2 /u/bananaboydean /u/harkins")
  if err != nil { fmt.Println(err) }
  r_eagles, err := r.AboutSubreddit("eagles")
  if err != nil { fmt.Println(err) }
  last := ""
  body := ""
  for {
    comments, err := r_eagles.Comments(&r.Session, 5, "", last)
    if err != nil { fmt.Println(err) }
    if len(comments) > 0 {
      last = comments[0].FullID
    }
    for _, comment := range comments {
      fmt.Println(comment.Body)
      body = strings.ToLower(comment.Body)
      if strings.Contains(body, "cowboys") {
        r.Reply(comment, "Fuck the cowboys")
        fmt.Println(comment.Body)
        fmt.Println("Replied")
      }
    }
    time.Sleep(time.Second*120)
  }
}
