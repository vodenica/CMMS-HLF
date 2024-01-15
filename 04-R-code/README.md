# The "R" code

The `R` code is used to plot the results of the experiments. The code is written in the `R` programming language and uses the `ggplot2` library to create the plots. The code is executed in the `RStudio` IDE.

## Average Latency

The average latency is calculated and plotted by the following R code:

```R
# Import required library
library(MASS)
library(tidyverse)

# Import CSV data
data <- read.csv("C:/[path-to-the-CSV-file]/[csv-file].csv")

str(data)
summary(data)
data

# Reshape the data to long format
data_long <- data %>%
  gather(variable, value, -transactionLoad)

# Define the desired order of the legend variable levels
legend_order <- c("10 Tx/sec", "50 Tx/sec", "100 Tx/sec", "200 Tx/sec", "500 Tx/sec")

# Create a new factor variable with the desired order
data_long$transactionLoad <- factor(data_long$transactionLoad, levels = legend_order)

# Define the desired order of the "variable" column
variable_order <- c("X10.Tx","X50.Tx", "X100.Tx", "X200.Tx", "X500.Tx")

# Convert the "variable" column to a factor with the desired order
data_long$variable <- factor(data_long$variable, levels = variable_order)

# Create a line plot
ggplot(data_long, aes(x = variable, y = value, group = transactionLoad , color = transactionLoad)) +
  geom_line(size = 1.5) +  # Increase the line thickness
  labs(title = NULL,
       x = "Number of Transactions [Tx]",
       y = "Average latency [sec]") +
  theme_minimal()+
  guides(color = guide_legend(title = NULL, 
                              override.aes = list(size = 1))) +  # Increase legend line size
  theme(legend.position = c(.05, .95),  # Position the legend inside the plot in the upper left corner
        legend.justification = c("left", "top"),
        legend.box.just = "right",
        legend.margin = margin(1, 1, 1, 1),
        plot.margin = margin(0.5, 0.5, 0.5, 0.5, "cm"),  # Add margins around the plot
        panel.border = element_rect(colour = "black", fill = NA, size = 1),  # Add a border around the plot
        axis.text.x = element_text(size = rel(2)),  # Increase x-axis text size by 100%
        axis.text.y = element_text(size = rel(2)),  # Increase y-axis text size by 100%
        axis.title.x = element_text(size = rel(2)),  # Increase x-axis title size by 100%
        axis.title.y = element_text(size = rel(2)),  # Increase y-axis title size by 100%
        legend.text = element_text(size = rel(2)),  # Increase legend text size by 100%
        plot.title = element_text(size = rel(2), hjust = 0.5)) +  # Increase main title size by 100% and center it
  scale_x_discrete(labels = c("10 Tx", "50 Tx", "100 Tx", "200 Tx", "500 Tx")) +
  expand_limits(y = 0)  # Include 0.0 on the Y-axis

```
## Throughput

The throughput is calculated and plotted by the following R code:

```R
# Import required library
library(MASS)
library(tidyverse)

# Import CSV data
data <- read.csv("C:/[path-to-the-CSV-file]/[csv-file].csv")

str(data)
summary(data)
data

# Reshape the data to long format
data_long <- data %>%
  gather(variable, value, -transactionLoad)

# Define the desired order of the legend variable levels
legend_order <- c("10 Tx/sec", "50 Tx/sec", "100 Tx/sec", "200 Tx/sec", "500 Tx/sec")

# Create a new factor variable with the desired order
data_long$transactionLoad <- factor(data_long$transactionLoad, levels = legend_order)

# Define the desired order of the "variable" column
variable_order <- c("X10.Tx", "X50.Tx", "X100.Tx", "X200.Tx", "X500.Tx")

# Convert the "variable" column to a factor with the desired order
data_long$variable <- factor(data_long$variable, levels = variable_order)

# Create a line plot
ggplot(data_long, aes(x = variable, y = value, group = transactionLoad , color = transactionLoad)) +
  geom_line(size = 1.5) +
  labs(title = NULL,
       x = "Number of Transactions [Tx]",
       #x = NULL,
       y = "Throughput [TPS]") +
  theme_minimal()+
  guides(color = guide_legend(title = NULL, 
                              override.aes = list(size = 1))) +  # Increase legend line size
  theme(legend.position = c(.05, .95),  # Position the legend inside the plot in the upper left corner
        legend.justification = c("left", "top"),
        legend.box.just = "right",
        legend.margin = margin(1, 1, 1, 1),
        plot.margin = margin(0.5, 0.5, 0.5, 0.5, "cm"),  # Add margins around the plot
        panel.border = element_rect(colour = "black", fill = NA, size = 1),  # Add a border around the plot
        axis.text.x = element_text(size = rel(2)),  # Increase x-axis text size by 100%
        axis.text.y = element_text(size = rel(2)),  # Increase y-axis text size by 100%
        axis.title.x = element_text(size = rel(2)),  # Increase x-axis title size by 100%
        axis.title.y = element_text(size = rel(2)),  # Increase y-axis title size by 100%
        legend.text = element_text(size = rel(2)),  # Increase legend text size by 100%
        plot.title = element_text(size = rel(2), hjust = 0.5)) +  # Increase main title size by 100% and center it
  scale_x_discrete(labels = c("10 Tx", "50 Tx", "100 Tx", "200 Tx", "500 Tx")) +
  expand_limits(y = 0)  # Include 0.0 on the Y-axis
```
# READ function throughput

The READ function throughput is calculated and plotted by the following R code:

```R
# Load necessary library
library(ggplot2)

# Read the CSV file
data <- read.csv("C:/[path-to-the-CSV-file]/[csv-file].csv")

# Calculate the statistics
stats <- data.frame(
  min = min(data$transactions),
  median = median(data$transactions),
  max = max(data$transactions)
)

# Create a boxplot for the specified column
p <- ggplot(data, aes(x = '', y = transactions)) + 
  geom_boxplot(outlier.shape = 1, outlier.size = 4, outlier.colour = "red", width = 0.25, colour = "blue", fill = "skyblue") +
  labs(y = "Throughput [TPS]") +
  theme_minimal() +  # Set the background to white
  theme(plot.margin = margin(0.5, 0.5, 0.5, 0.5, "cm"),  # Add margins around the plot
        panel.border = element_rect(colour = "black", fill = NA, size = 1),  # Add a border around the plot
        axis.text.x = element_blank(),  # Remove x-axis text
        axis.ticks.x = element_blank(),  # Remove x-axis ticks
        axis.text.y = element_text(size = rel(2)),  # Increase y-axis text size by 100%
        axis.title.y = element_text(size = rel(2)))  # Increase y-axis title size by 100%

# Add small circles at the maximum and minimum values
p <- p + geom_point(data = stats, aes(x = '', y = min), size = 3, shape = 16, color = "blue") +  # Add a circle at the minimum value
  geom_point(data = stats, aes(x = '', y = max), size = 3, shape = 16, color = "blue")  # Add a circle at the maximum value

# Print the min, max, and median values on the plot
p <- p + annotate("text", x = 0.5, y = stats$min, label = paste("Min:", stats$min), hjust = -0.1, size = 5, color = "blue") +
  annotate("text", x = 0.5, y = stats$median, label = paste("Median:", stats$median), hjust = -0.1, size = 5, color = "blue") +
  annotate("text", x = 0.5, y = stats$max, label = paste("Max:", stats$max), hjust = -0.1, size = 5, color = "blue")

# Print the plot
print(p)
```
## Plotting the results

The plots are saved as PNG files and can be found in the `05-Plots/mte` folder. The plots are also shown below only for the `maintenance` chaincode.

