package main

import (
	"fmt"
	"os"

	"github.com/godbus/dbus/v5"
)

func main() {
	// Connect to the system D-Bus.
	conn, err := dbus.ConnectSystemBus()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to connect to system bus: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	// Create a D-Bus object to interact with the D-Bus service itself.
	// We'll call a method on the org.freedesktop.DBus service at the
	// /org/freedesktop/DBus object path, using the
	// org.freedesktop.DBus.ListNames interface.
	obj := conn.Object("org.freedesktop.DBus", "/org/freedesktop/DBus")

	// Call the "ListNames" method. This method returns a list of all
	// currently registered names on the D-Bus.
	var names []string
	err = obj.Call("org.freedesktop.DBus.ListNames", 0).Store(&names)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to call method: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Registered D-Bus names:")
	for _, name := range names {
		fmt.Println("- ", name)
	}
}
