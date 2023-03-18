antlr.bat -listener -no-visitor -Dlanguage=Go -package preludiocompiler preludio.g4

Move-Item -force .\preludio_* .\core\compiler\
Move-Item -force .\preludio*.interp .\core\compiler\
Move-Item -force .\preludio*.tokens .\core\compiler\