library(tidyverse)

a <- tibble(
    xa = c(1, 2, 3, 4, 5, 6, 7, 8, 9, 10),
    y = c("a", "a", "c", "c", "e", "f", "f", "h", "i", "j"),
    z = c("id1", "id1", "id2", "id2", "id2", "id5", "id6", "id7", "id8", "id8"),
)

b <- tibble(
    xb = c(5, 6, 7, 8, 9, 10, 11, 12, 13, 14),
    y = c("c", "c", "c", "d", "d", "d", "e", "e", "e", "f"),
    z = c("id2", "id2", "id5", "id6", "id7", "id8", "id9", "id10", "id11", "id11"),
)


