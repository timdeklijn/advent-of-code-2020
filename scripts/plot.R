library("ggplot2")

# Load data
table <- read.table("table", header=TRUE)

table$dt2 <- table$dt2 - table$dt1
sum <- table$dt2 + table$dt1
max_y <- max(sum)
max_y <- max_y + (0.05*max_y)

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
        size=0.3,
        stat="identity",
        position=position_stack(reverse=TRUE)
        ) +
    # Set Theme, both color theme and add a border around the legend
    theme_classic(base_size=20) + 
    theme(
        # Legend colors
        legend.background=element_rect(fill="#0f0f23"),
        legend.text=element_text(color="#009900"),
        legend.title=element_text(color="#00cc00"),
        # Set backgrounds
        plot.background=element_rect(fill="#0f0f23"),
        panel.background=element_rect(fill="#0f0f23"),
        # Axis colors
        axis.text=element_text(color="#009900"),
        axis.line=element_line(color="#cccccc"),
        axis.title=element_text(color="#00cc00"),
        axis.ticks=element_line(color="#cccccc"),
        # Plot title
        plot.title=element_text(color="#00cc00"),
        plot.subtitle=element_text(color="#009900"),
    ) +
    # Set Labels
    labs(
        title="Advent of Code 2020",
        subtitle="Solution time puzzle 1 and puzzle 2",
        x="day",
        y="Solve Time (Hours)"
    ) +
    # Move axis origins
    expand_limits(x = 0.5, y = 0) +
    # Reset X-ticks
    scale_x_continuous(breaks=seq(1, max(table$day), 1), labels=sprintf("%d", seq(1, max(table$day), 1))) +
    # Move x to y=0
    scale_y_continuous(expand = c(0, 0), limits=c(0, max_y)) +
    scale_fill_manual(values=c("#009900", "#ffff66")) + 
    png("tst.png", width=1025, height=512)

