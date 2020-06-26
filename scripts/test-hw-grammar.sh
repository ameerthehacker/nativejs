source scripts/generate-parser.sh

ROOT=$(pwd)
echo $ROOT
cd internal/parser
antlr NJSLexer.g4 NJSParser.g4
javac *.java
grun NJS program -gui $ROOT/test/hello-world.njs