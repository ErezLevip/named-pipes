# Named Pipes

named pipes is a package meant to simplify the use of named pipes in GO.

## How to install
     go get github.com/erezLevip/named-pipes

## Example 
     p, err := pkg.NewPipe("main.p")
     if err != nil {
     	log.Fatal(err)
     }
     go writeTests(p)
     for v := range p.Listen('\n') {
     	if err == nil {
     		fmt.Print("value from channel:" + string(v))
     	}
     }