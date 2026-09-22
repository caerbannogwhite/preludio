# Regenerates the ANTLR lexer/parser into src/bytefeeder/.
#
# The committed files were generated with ANTLR 4.11.1 (see the
# "Code generated ... by ANTLR 4.11.1" header in src/bytefeeder/preludio_lexer.go).
# Anyone editing preludioLexer.g4 / preludioParser.g4 must use that exact
# toolchain so the checked-in parser stays reproducible:
#   https://www.antlr.org/download/antlr-4.11.1-complete.jar
# Expose it on PATH as antlr.bat (Windows) or antlr.

$antlr = Get-Command antlr.bat -ErrorAction SilentlyContinue
if (-not $antlr) {
    $antlr = Get-Command antlr -ErrorAction SilentlyContinue
}
if (-not $antlr) {
    Write-Error "antlr/antlr.bat not found on PATH. Install ANTLR 4.11.1 from https://www.antlr.org/download/antlr-4.11.1-complete.jar and retry."
    exit 1
}

& $antlr.Source -listener -no-visitor -Dlanguage=Go -package bytefeeder preludioLexer.g4 preludioParser.g4
if ($LASTEXITCODE -ne 0) {
    exit $LASTEXITCODE
}

Move-Item -force .\preludio*.go .\src\bytefeeder\
Move-Item -force .\preludio*.interp .\src\bytefeeder\
Move-Item -force .\preludio*.tokens .\src\bytefeeder\
