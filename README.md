# go-plays

Another study project to copy site, the entry point should trigger a service to build local storage, 
And Theme => Generate online
And topic => should have a way to extract the serivces from local service

storage/ for storing imgs?
controllers/ controll business logic
modules 

思路：
获取主题

type Theme struct {
   TheTitle string
   TheUrl string
   TheLimit int
}

type Subject struct {
   SubTitle string
   SubUrl string
   SubLimit int
}

type Image struct {
   ImageName string
   ImageUrl string
}