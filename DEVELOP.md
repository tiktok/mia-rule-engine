## ANTLR 4 Setup
- Install antlr by using `brew`
``` bash
brew install antlr
```
- Re-generate ANTLR codes
```bash
make gen-antlr
```

## Rete Network Virtualization
- Install graphviz by using `brew`
``` bash
brew install graphviz
```
- Copy `network.Virtualize()` output to `rete_network.dot` file and run the command below to generate the png picture  
```bash
dot -Tpng -o .assets/rete_network.png .assets/rete_network.dot
```

## Performance Analyze with `pprof`
- Add the following code snippet to benchmark testing
```Golang
// create pprof output file
f, err := os.Create("/tmp/cpu_profile.out")
if err != nil {
    b.Fatal(err)
}
defer f.Close()

// start pprof
pprof.StartCPUProfile(f)
defer pprof.StopCPUProfile()
```
- Run the command to enter command line
```bash
go tool pprof /tmp/cpu_profile.out
```
- Enter WebUI
```bash
(pprof) web
```

