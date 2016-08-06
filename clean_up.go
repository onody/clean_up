package main

import (
  "fmt"
  "os"
  "io/ioutil"
  "path/filepath"

  "github.com/urfave/cli"
)

func main() {
  app := cli.NewApp()
  app.Name = "clean_up"
  app.Usage = "deletes garbage files"
  app.Version = Version
  app.Author = "onody"
  app.Email = "onodera212@gmail.com"
  app.Action = func(c *cli.Context) error {
    path := c.Args().Get(0)
    doCleanUp(path, path)
    return nil
  }
  app.Run(os.Args)
}

func doCleanUp(rootPath string, searchPath string) {
  fis, err := ioutil.ReadDir(searchPath)
  if err != nil {
    panic(err)
  }

  for _, fi := range fis {
    fullPath := filepath.Join(searchPath, fi.Name())

    delete(fullPath, fi.Name())
    if fi.IsDir() {
      doCleanUp(rootPath, fullPath)
    }
  }
}

func delete(fullPath string, name string){
  for _,g := range targetFiles {
    if(name == g){
      if err := os.Remove(fullPath); err != nil {
        fmt.Println(err)
      }else{
        fmt.Println("deleted!! -> " + fullPath)
      }
    }
  }
}
