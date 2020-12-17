library("ggplot2")

# Load data
table <- read.table("table", header=TRUE)

# Split data, add puzzle_number, and merge back
tmp1 <- table[, c("day", "dt1")]
names(tmp1)[names(tmp1) == "dt1"] <- "dt"
tmp1$puzzle_number <- rep(c("1"), times=nrow(tmp1))

tmp2 <- table[, c("day", "dt2")]
names(tmp2)[names(tmp2) == "dt2"] <- "dt"
tmp2$puzzle_number <- rep(c("2"), times=nrow(tmp2))

table <- rbind(tmp1, tmp2)

# Add data
ggplot(table) + 
    # Bar Chart
    geom_bar(
        aes(x=day, y=dt, fill=puzzle_number),
        colour="black",
        stat="identity",
        position=position_dodge()
        ) +
    # Set Theme, both color theme and add a border around the legend
    theme_bw() +
    theme(legend.background = element_rect(colour = 'black', fill = 'white', linetype='solid')) +
    # Set Labels
    labs(
        title="Advent of Code 2020",
        subtitle="Compare solution time puzzle 1 and puzzle 2",
        x="day",
        y="Solve Time (minutes)"
        ) +
    # Move axis origins
    expand_limits(x = 0.5, y = 0) +
    # Reset X-ticks
    scale_x_continuous(breaks=seq(1, max(table$day), 1), labels=sprintf("%d", seq(1, max(table$day), 1))) +
    # Move x to y=0
    scale_y_continuous(expand = c(0, 0)) +
    png("tst.png")