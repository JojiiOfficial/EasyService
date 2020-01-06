# EasySerivce
A cli tool to create a systemd service easily from cli


# Install
run
```go
go get
go build -o ezservice
```
to build it.
<br>Then you can run `sudo mv ezservice /bin/ezservice` if you want.<br>

# Usage

<b>Create a service</b><br>
`sudo ezservice create -N <name> -F /path/to/execfile -G root -U root -s -e`<br>
`-N` Name<br>
`-D` Description<br>
`-F` ExecStart file<br>
`-C` ExecStart command<br>
`-G` Group<br>
`-U` User<br>
`-s` Start service afterwards<br>
`-e` Enable service afterwards<br>
`-y` Yes (skip confirm message)<br>

<b>Delete a service</b><br>
`sudo ezservice delete -N <name> -y`<br>
`-N` Name<br>
`-y` Yes (skip confirm message)<br>
