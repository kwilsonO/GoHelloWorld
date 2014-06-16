GC = go  
BC = build 
TARG = HelloWorld 
BTARG = HelloWorld.go
PACKAGE = main

all:
	make clean
	mkdir $(PACKAGE)
	cp $(BTARG) $(PACKAGE)/
	make proj
	cp $(TARG) $(PACKAGE)/
	rm $(TARG)        

proj:
	$(GC) $(BC) $(PACKAGE)/$(BTARG)


clean:
	rm -rf  $(PACKAGE)
