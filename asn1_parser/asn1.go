package asn1_parser

//go:generate wget -c https://www.antlr.org/download/antlr-4.10.1-complete.jar
//go:generate java -jar antlr-4.10.1-complete.jar -o ./ -package asn1_parser -listener -visitor -Dlanguage=Go asn1.g4
