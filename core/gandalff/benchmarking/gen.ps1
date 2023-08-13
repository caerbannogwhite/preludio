# With 0% NAs
Rscript.exe groupby-datagen.R 1e4 1e2 0 0
Rscript.exe groupby-datagen.R 1e5 1e2 0 0
Rscript.exe groupby-datagen.R 1e6 1e2 0 0
Rscript.exe groupby-datagen.R 1e7 1e2 0 0

# With 10% NAs
Rscript.exe groupby-datagen.R 1e4 1e2 10 0
Rscript.exe groupby-datagen.R 1e5 1e2 10 0
Rscript.exe groupby-datagen.R 1e6 1e2 10 0
Rscript.exe groupby-datagen.R 1e7 1e2 10 0

Move-Item -Force G1_* ..\testdata\