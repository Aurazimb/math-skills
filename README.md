# Statistics Calculator

This simple program reads data from a file and calculates various statistics on the population data provided in the file. The program can be used to compute the average, median, variance, and standard deviation of the data.

## Usage

You can run this program using a command-line interface. If your project is written in Go, you can use the following command:

```shell
$ go run your-program.go data.txt
```

Replace `your-program.go` with the name of your Go program and `data.txt` with the path to the file containing the population data.

## Input Format

The input file should contain one value per line, representing the statistical population. For example:

```
189
113
121
114
145
110
...
```

## Output Format

After reading the data from the file, the program will execute the following calculations and print the results as rounded integers:

- **Average:** The program will calculate and print the average of the data.
- **Median:** The program will determine and print the median of the data.
- **Variance:** The program will compute and print the variance of the data.
- **Standard Deviation:** The program will calculate and print the standard deviation of the data.

Here is an example of how the output might look:

```shell
Average: 35
Median: 4
Variance: 5
Standard Deviation: 65
```

## Note

The program will round the calculated values to the nearest integer as per your requirements.

Feel free to use and modify this program to suit your needs. Enjoy your statistical analysis!