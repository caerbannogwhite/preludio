# DEBUG SCRIPT
# Remove-Item .\.antlr\*
# antlr.bat preludioLexer.g4 preludioParser.g4 -o .\.antlr
# javac.exe .\.antlr\*.java

antlr.bat -listener -no-visitor -Dlanguage=Go -package bytefeeder preludioLexer.g4 preludioParser.g4

Move-Item -force .\preludio*.go .\core\bytefeeder\
Move-Item -force .\preludio*.interp .\core\bytefeeder\
Move-Item -force .\preludio*.tokens .\core\bytefeeder\