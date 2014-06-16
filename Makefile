GC = go  
BC = build 
TARG = HelloWorld 
BTARG = HelloWorld.go
PACKAGE = package

all:
	make package
package:
	make clean
	mkdir $(PACKAGE)
	cp -r web $(PACKAGE)/
	cp $(BTARG) $(PACKAGE)/
	make proj
	cp $(TARG) $(PACKAGE)/
	rm $(TARG)        

proj:
	$(GC) $(BC) $(PACKAGE)/$(BTARG)


clean:
	rm -rf  $(PACKAGE)
