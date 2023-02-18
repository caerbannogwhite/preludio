antlr.bat -listener -no-visitor -Dlanguage=Go -package compiler preludio.g4

Move-Item -force .\preludio_* .\preludio\compiler\
Move-Item -force .\preludio*.interp .\preludio\compiler\
Move-Item -force .\preludio*.tokens .\preludio\compiler\