args <- commandArgs(trailingOnly = TRUE)

print(args)

options(shiny.launch.browser = FALSE)

shiny::runExample("01_hello", port = as.numeric(args))