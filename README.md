
![](media/logo_med.png)

# Preludio

### A PRQL based data transformation language
Preludio is a data transformation language based on PRQL. It is a language that allows you to transform and
manipulate data in a simple and intuitive way, batteries included.

No libraries or external dependencies are required to run the language.

### Example
This is a simple example of what you can already do with Preludio.
It reads a CSV file, derives two new columns, selects some columns and writes the result to a new CSV file.

```
readCSV "tests\\Cars.csv" delimiter: ";" header:true
derive [
  cylTimes2 = Cylinders * 2,
  Car_Origin = Car + " - " + Origin
]
take 5
select [Car_Origin, MPG, cylTimes2]
writeCSV "tests\\Cars1.csv"
```

![](media/repl_example.gif)

### Features
-  [x] Read and write CSV files
-  [x] Derive new columns
-  [x] Select columns

### Installation
To run it, you need to have [Go](https://golang.org/doc/install) installed.
Once you have Go, you can clone this repository.

To run the program, you can use the following command:
```bash
go run .
```

### Future Features
- [ ] Add support for more data types
- [ ] Move to Gandalff library
- [ ] Add more transformations functions (group by, etc.)
- [ ] Add statistical functions
- [ ] Add support for Excel files
- [ ] Add support for XPT files
- [ ] Add support for SAS7BDAT files
- [ ] Database connections (SQL, MongoDB, etc.)
- [ ] VS Code extension

In case the language becomes quite successful, I will consider adding:

- [ ] Text editor/Ide (in browser and app)
- [ ] Plots (https://github.com/gonum/plot)
- [ ] Interactive plots and tables
- [ ] Integration with other languages (Python, R, etc.)
- [ ] Integration with other tools (Jupyter, etc.)
- [ ] Integration with OpenAI (https://openai.com/blog/openai-api/), ie. image to table

### Contributing
If you want to contribute to this project, you can do so by forking the repository and submitting a pull request.

### Developers
If the grammar is changed, the parser must be regenerated. To do this, run the following command:

(on Windows)
```
make.ps1
```

### Log
 - **21 / 03 / 2023** First publishing of the repository. Many things are still not working.
