package main

import (
    "fmt"
    "os"
    "strconv"
  "github.com/codegangsta/cli"
)

func main() {
    app := cli.NewApp()
    app.Name = "naval_fate"
    app.Commands = []cli.Command{
        {
            Name: "ship",
            Subcommands: []cli.Command{
                {
                    Name:   "new",
                    Usage:  "NAME...",
                    Action: newShips,
                },
                {
                    Name:  "move",
                    Usage: "NAME X Y",
                    Flags: []cli.Flag{
                        cli.IntFlag{
                            Name:  "speed",
                            Value: 10,
                            Usage: "Speed in knots",
                        },
                    },
                    Action: moveShip,
                },
            },
        },
        {
            Name: "mine",
            Subcommands: []cli.Command{
                {
                    Name:  "set",
                    Usage: "X Y [--moored|--drifting]",
                    Flags: []cli.Flag{
                        cli.BoolFlag{
                            Name:  "moored",
                            Usage: "Moored (anchored) mine",
                        },
                        cli.BoolFlag{
                            Name:  "drifting",
                            Usage: "Drifting mine",
                        },
                    },
                    Action: setMine,
                },
            },
        },
    }

    app.Run(os.Args)
}

func newShips(c *cli.Context) {
    if len(c.Args()) == 0 {
        fmt.Printf("Incorrect usage\n")

        return
    }
    fmt.Printf("create ships %#v\n", c.Args())
}

func moveShip(c *cli.Context) {
    if len(c.Args()) != 3 {
        fmt.Printf("Incorrect usage\n")
        return
    }
    name := c.Args()[0]
    x, err := strconv.Atoi(c.Args()[1])
    if err != nil {
        fmt.Printf("Incorrect usage\n")
        return
    }
    y, err := strconv.Atoi(c.Args()[2])
    if err != nil {
        fmt.Printf("Incorrect usage\n")
        return
    }
    speed := c.Int("speed")
    fmt.Printf("move ship named %v to %d:%d with speed %d\n", name, x, y, speed)
}

func setMine(c *cli.Context) {
    if len(c.Args()) != 2 {
        fmt.Printf("Incorrect usage\n")
        return
    }
    x, err := strconv.Atoi(c.Args()[0])
    if err != nil {
        fmt.Printf("Incorrect usage\n")
        return
    }
    y, err := strconv.Atoi(c.Args()[1])
    if err != nil {
        fmt.Printf("Incorrect usage\n")
        return
    }
    if c.Bool("moored") && c.Bool("drifting") {
        fmt.Printf("Incorrect usage\n")
        return
    }
    s := "moored"
    switch {
    case c.Bool("moored"):
        s = "moored"
    case c.Bool("drifting"):
        s = "drifting"
    }
    fmt.Printf("set a %s mine in %d:%d\n", s, x, y)
}
