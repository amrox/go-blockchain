{ pkgs ? import <nixpkgs> {} }:
with pkgs;

mkShell {
    # nativeBuildInputs is usually what you want -- tools you need to run
    nativeBuildInputs = [ 
    	#buildPackages.go_1_17
    	go_1_17
	gotools
    	gopls
    	go-outline
    	gocode
    	gopkgs
    	gocode-gomod
    	godef
    	golint
    ];
}
