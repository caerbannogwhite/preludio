#!/usr/bin/env Rscript

cat("# groupby-dplyr.R\n")

source("./_helpers/helpers.R")

stopifnot(requireNamespace(c("bit64","data.table"), quietly=TRUE)) # used in chk to sum numeric columns and data loading
.libPaths("./dplyr/r-dplyr") # tidyverse/dplyr#4641
suppressPackageStartupMessages(library("dplyr", lib.loc="./dplyr/r-dplyr", warn.conflicts=FALSE))
ver = packageVersion("dplyr")
git = "" # uses stable version now #124
task = "groupby"
solution = "dplyr"
fun = "group_by"
cache = TRUE
on_disk = FALSE

data_name = Sys.getenv("SRC_DATANAME")
src_grp = file.path("data", paste(data_name, "csv", sep="."))
cat(sprintf("loading dataset %s\n", data_name))

x = as_tibble(data.table::fread(src_grp, showProgress=FALSE, stringsAsFactors=TRUE, na.strings="", data.table=FALSE))
print(nrow(x))

task_init = proc.time()[["elapsed"]]
cat("grouping...\n")

question = "sum v1 by id1" # q1
t = system.time(print(dim(ans<-x %>% group_by(id1) %>% summarise(v1=sum(v1, na.rm=TRUE)))))[["elapsed"]]
m = memory_usage()
chkt = system.time(chk<-summarise(ungroup(ans), v1=sum(bit64::as.integer64(v1))))[["elapsed"]]
write.log(run=1L, task=task, data=data_name, in_rows=nrow(x), question=question, out_rows=nrow(ans), out_cols=ncol(ans), solution=solution, version=ver, git=git, fun=fun, time_sec=t, mem_gb=m, cache=cache, chk=make_chk(chk), chk_time_sec=chkt, on_disk=on_disk)
rm(ans)
t = system.time(print(dim(ans<-x %>% group_by(id1) %>% summarise(v1=sum(v1, na.rm=TRUE)))))[["elapsed"]]
m = memory_usage()
chkt = system.time(chk<-summarise(ungroup(ans), v1=sum(bit64::as.integer64(v1))))[["elapsed"]]
write.log(run=2L, task=task, data=data_name, in_rows=nrow(x), question=question, out_rows=nrow(ans), out_cols=ncol(ans), solution=solution, version=ver, git=git, fun=fun, time_sec=t, mem_gb=m, cache=cache, chk=make_chk(chk), chk_time_sec=chkt, on_disk=on_disk)
print(head(ans, 3))
print(tail(ans, 3))
rm(ans)

question = "sum v1 by id1:id2" # q2
t = system.time(print(dim(ans<-x %>% group_by(id1, id2) %>% summarise(v1=sum(v1, na.rm=TRUE)))))[["elapsed"]]
m = memory_usage()
chkt = system.time(chk<-summarise(ungroup(ans), v1=sum(bit64::as.integer64(v1))))[["elapsed"]]
write.log(run=1L, task=task, data=data_name, in_rows=nrow(x), question=question, out_rows=nrow(ans), out_cols=ncol(ans), solution=solution, version=ver, git=git, fun=fun, time_sec=t, mem_gb=m, cache=cache, chk=make_chk(chk), chk_time_sec=chkt, on_disk=on_disk)
rm(ans)
t = system.time(print(dim(ans<-x %>% group_by(id1, id2) %>% summarise(v1=sum(v1, na.rm=TRUE)))))[["elapsed"]]
m = memory_usage()
chkt = system.time(chk<-summarise(ungroup(ans), v1=sum(bit64::as.integer64(v1))))[["elapsed"]]
write.log(run=2L, task=task, data=data_name, in_rows=nrow(x), question=question, out_rows=nrow(ans), out_cols=ncol(ans), solution=solution, version=ver, git=git, fun=fun, time_sec=t, mem_gb=m, cache=cache, chk=make_chk(chk), chk_time_sec=chkt, on_disk=on_disk)
print(head(ans, 3))
print(tail(ans, 3))
rm(ans)

question = "sum v1 mean v3 by id3" # q3
t = system.time(print(dim(ans<-x %>% group_by(id3) %>% summarise(v1=sum(v1, na.rm=TRUE), v3=mean(v3, na.rm=TRUE)))))[["elapsed"]]
m = memory_usage()
chkt = system.time(chk<-summarise(ungroup(ans), v1=sum(bit64::as.integer64(v1)), v3=sum(v3)))[["elapsed"]]
write.log(run=1L, task=task, data=data_name, in_rows=nrow(x), question=question, out_rows=nrow(ans), out_cols=ncol(ans), solution=solution, version=ver, git=git, fun=fun, time_sec=t, mem_gb=m, cache=cache, chk=make_chk(chk), chk_time_sec=chkt, on_disk=on_disk)
rm(ans)
t = system.time(print(dim(ans<-x %>% group_by(id3) %>% summarise(v1=sum(v1, na.rm=TRUE), v3=mean(v3, na.rm=TRUE)))))[["elapsed"]]
m = memory_usage()
chkt = system.time(chk<-summarise(ungroup(ans), v1=sum(bit64::as.integer64(v1)), v3=sum(v3)))[["elapsed"]]
write.log(run=2L, task=task, data=data_name, in_rows=nrow(x), question=question, out_rows=nrow(ans), out_cols=ncol(ans), solution=solution, version=ver, git=git, fun=fun, time_sec=t, mem_gb=m, cache=cache, chk=make_chk(chk), chk_time_sec=chkt, on_disk=on_disk)
print(head(ans, 3))
print(tail(ans, 3))
rm(ans)

question = "mean v1:v3 by id4" # q4
t = system.time(print(dim(ans<-x %>% group_by(id4) %>% summarise_at(.funs="mean", .vars=c("v1","v2","v3"), na.rm=TRUE))))[["elapsed"]]
m = memory_usage()
chkt = system.time(chk<-summarise(ungroup(ans), v1=sum(v1), v2=sum(v2), v3=sum(v3)))[["elapsed"]]
write.log(run=1L, task=task, data=data_name, in_rows=nrow(x), question=question, out_rows=nrow(ans), out_cols=ncol(ans), solution=solution, version=ver, git=git, fun=fun, time_sec=t, mem_gb=m, cache=cache, chk=make_chk(chk), chk_time_sec=chkt, on_disk=on_disk)
rm(ans)
t = system.time(print(dim(ans<-x %>% group_by(id4) %>% summarise_at(.funs="mean", .vars=c("v1","v2","v3"), na.rm=TRUE))))[["elapsed"]]
m = memory_usage()
chkt = system.time(chk<-summarise(ungroup(ans), v1=sum(v1), v2=sum(v2), v3=sum(v3)))[["elapsed"]]
write.log(run=2L, task=task, data=data_name, in_rows=nrow(x), question=question, out_rows=nrow(ans), out_cols=ncol(ans), solution=solution, version=ver, git=git, fun=fun, time_sec=t, mem_gb=m, cache=cache, chk=make_chk(chk), chk_time_sec=chkt, on_disk=on_disk)
print(head(ans, 3))
print(tail(ans, 3))
rm(ans)

question = "sum v1:v3 by id6" # q5
t = system.time(print(dim(ans<-x %>% group_by(id6) %>% summarise_at(.funs="sum", .vars=c("v1","v2","v3"), na.rm=TRUE))))[["elapsed"]]
m = memory_usage()
chkt = system.time(chk<-summarise(ungroup(ans), v1=sum(bit64::as.integer64(v1)), v2=sum(bit64::as.integer64(v2)), v3=sum(v3)))[["elapsed"]]
write.log(run=1L, task=task, data=data_name, in_rows=nrow(x), question=question, out_rows=nrow(ans), out_cols=ncol(ans), solution=solution, version=ver, git=git, fun=fun, time_sec=t, mem_gb=m, cache=cache, chk=make_chk(chk), chk_time_sec=chkt, on_disk=on_disk)
rm(ans)
t = system.time(print(dim(ans<-x %>% group_by(id6) %>% summarise_at(.funs="sum", .vars=c("v1","v2","v3"), na.rm=TRUE))))[["elapsed"]]
m = memory_usage()
chkt = system.time(chk<-summarise(ungroup(ans), v1=sum(bit64::as.integer64(v1)), v2=sum(bit64::as.integer64(v2)), v3=sum(v3)))[["elapsed"]]
write.log(run=2L, task=task, data=data_name, in_rows=nrow(x), question=question, out_rows=nrow(ans), out_cols=ncol(ans), solution=solution, version=ver, git=git, fun=fun, time_sec=t, mem_gb=m, cache=cache, chk=make_chk(chk), chk_time_sec=chkt, on_disk=on_disk)
print(head(ans, 3))
print(tail(ans, 3))
rm(ans)

question = "median v3 sd v3 by id4 id5" # q6
t = system.time(print(dim(ans<-x %>% group_by(id4, id5) %>% summarise(median_v3=median(v3, na.rm=TRUE), sd_v3=sd(v3, na.rm=TRUE)))))[["elapsed"]]
m = memory_usage()
chkt = system.time(chk<-summarise(ungroup(ans), median_v3=sum(median_v3), sd_v3=sum(sd_v3)))[["elapsed"]]
write.log(run=1L, task=task, data=data_name, in_rows=nrow(x), question=question, out_rows=nrow(ans), out_cols=ncol(ans), solution=solution, version=ver, git=git, fun=fun, time_sec=t, mem_gb=m, cache=cache, chk=make_chk(chk), chk_time_sec=chkt, on_disk=on_disk)
rm(ans)
t = system.time(print(dim(ans<-x %>% group_by(id4, id5) %>% summarise(median_v3=median(v3, na.rm=TRUE), sd_v3=sd(v3, na.rm=TRUE)))))[["elapsed"]]
m = memory_usage()
chkt = system.time(chk<-summarise(ungroup(ans), median_v3=sum(median_v3), sd_v3=sum(sd_v3)))[["elapsed"]]
write.log(run=2L, task=task, data=data_name, in_rows=nrow(x), question=question, out_rows=nrow(ans), out_cols=ncol(ans), solution=solution, version=ver, git=git, fun=fun, time_sec=t, mem_gb=m, cache=cache, chk=make_chk(chk), chk_time_sec=chkt, on_disk=on_disk)
print(head(ans, 3))
print(tail(ans, 3))
rm(ans)

question = "max v1 - min v2 by id3" # q7
t = system.time(print(dim(ans<-x %>% group_by(id3) %>% summarise(range_v1_v2=max(v1, na.rm=TRUE)-min(v2, na.rm=TRUE)))))[["elapsed"]]
m = memory_usage()
chkt = system.time(chk<-summarise(ungroup(ans), range_v1_v2=sum(bit64::as.integer64(range_v1_v2))))[["elapsed"]]
write.log(run=1L, task=task, data=data_name, in_rows=nrow(x), question=question, out_rows=nrow(ans), out_cols=ncol(ans), solution=solution, version=ver, git=git, fun=fun, time_sec=t, mem_gb=m, cache=cache, chk=make_chk(chk), chk_time_sec=chkt, on_disk=on_disk)
rm(ans)
t = system.time(print(dim(ans<-x %>% group_by(id3) %>% summarise(range_v1_v2=max(v1, na.rm=TRUE)-min(v2, na.rm=TRUE)))))[["elapsed"]]
m = memory_usage()
chkt = system.time(chk<-summarise(ungroup(ans), range_v1_v2=sum(bit64::as.integer64(range_v1_v2))))[["elapsed"]]
write.log(run=2L, task=task, data=data_name, in_rows=nrow(x), question=question, out_rows=nrow(ans), out_cols=ncol(ans), solution=solution, version=ver, git=git, fun=fun, time_sec=t, mem_gb=m, cache=cache, chk=make_chk(chk), chk_time_sec=chkt, on_disk=on_disk)
print(head(ans, 3))
print(tail(ans, 3))
rm(ans)

question = "largest two v3 by id6" # q8
t = system.time(print(dim(ans<-x %>% select(id6, largest2_v3=v3) %>% filter(!is.na(largest2_v3)) %>% arrange(desc(largest2_v3)) %>% group_by(id6) %>% filter(row_number() <= 2L))))[["elapsed"]]
m = memory_usage()
chkt = system.time(chk<-summarise(ungroup(ans), largest2_v3=sum(largest2_v3)))[["elapsed"]]
write.log(run=1L, task=task, data=data_name, in_rows=nrow(x), question=question, out_rows=nrow(ans), out_cols=ncol(ans), solution=solution, version=ver, git=git, fun=fun, time_sec=t, mem_gb=m, cache=cache, chk=make_chk(chk), chk_time_sec=chkt, on_disk=on_disk)
rm(ans)
t = system.time(print(dim(ans<-x %>% select(id6, largest2_v3=v3) %>% filter(!is.na(largest2_v3)) %>% arrange(desc(largest2_v3)) %>% group_by(id6) %>% filter(row_number() <= 2L))))[["elapsed"]]
m = memory_usage()
chkt = system.time(chk<-summarise(ungroup(ans), largest2_v3=sum(largest2_v3)))[["elapsed"]]
write.log(run=2L, task=task, data=data_name, in_rows=nrow(x), question=question, out_rows=nrow(ans), out_cols=ncol(ans), solution=solution, version=ver, git=git, fun=fun, time_sec=t, mem_gb=m, cache=cache, chk=make_chk(chk), chk_time_sec=chkt, on_disk=on_disk)
print(head(ans, 3))
print(tail(ans, 3))
rm(ans)

question = "regression v1 v2 by id2 id4" # q9
t = system.time(print(dim(ans<-x %>% group_by(id2, id4) %>% summarise(r2=cor(v1, v2, use="na.or.complete")^2))))[["elapsed"]]
m = memory_usage()
chkt = system.time(chk<-summarise(ungroup(ans), r2=sum(r2)))[["elapsed"]]
write.log(run=1L, task=task, data=data_name, in_rows=nrow(x), question=question, out_rows=nrow(ans), out_cols=ncol(ans), solution=solution, version=ver, git=git, fun=fun, time_sec=t, mem_gb=m, cache=cache, chk=make_chk(chk), chk_time_sec=chkt, on_disk=on_disk)
rm(ans)
t = system.time(print(dim(ans<-x %>% group_by(id2, id4) %>% summarise(r2=cor(v1, v2, use="na.or.complete")^2))))[["elapsed"]]
m = memory_usage()
chkt = system.time(chk<-summarise(ungroup(ans), r2=sum(r2)))[["elapsed"]]
write.log(run=2L, task=task, data=data_name, in_rows=nrow(x), question=question, out_rows=nrow(ans), out_cols=ncol(ans), solution=solution, version=ver, git=git, fun=fun, time_sec=t, mem_gb=m, cache=cache, chk=make_chk(chk), chk_time_sec=chkt, on_disk=on_disk)
print(head(ans, 3))
print(tail(ans, 3))
rm(ans)

question = "sum v3 count by id1:id6" # q10
t = system.time(print(dim(ans<-x %>% group_by(id1, id2, id3, id4, id5, id6) %>% summarise(v3=sum(v3, na.rm=TRUE), count=n()))))[["elapsed"]]
m = memory_usage()
chkt = system.time(chk<-summarise(ungroup(ans), v3=sum(v3), count=sum(bit64::as.integer64(count))))[["elapsed"]]
write.log(run=1L, task=task, data=data_name, in_rows=nrow(x), question=question, out_rows=nrow(ans), out_cols=ncol(ans), solution=solution, version=ver, git=git, fun=fun, time_sec=t, mem_gb=m, cache=cache, chk=make_chk(chk), chk_time_sec=chkt, on_disk=on_disk)
rm(ans)
t = system.time(print(dim(ans<-x %>% group_by(id1, id2, id3, id4, id5, id6) %>% summarise(v3=sum(v3, na.rm=TRUE), count=n()))))[["elapsed"]]
m = memory_usage()
chkt = system.time(chk<-summarise(ungroup(ans), v3=sum(v3), count=sum(bit64::as.integer64(count))))[["elapsed"]]
write.log(run=2L, task=task, data=data_name, in_rows=nrow(x), question=question, out_rows=nrow(ans), out_cols=ncol(ans), solution=solution, version=ver, git=git, fun=fun, time_sec=t, mem_gb=m, cache=cache, chk=make_chk(chk), chk_time_sec=chkt, on_disk=on_disk)
print(head(ans, 3))
print(tail(ans, 3))
rm(ans)

cat(sprintf("grouping finished, took %.0fs\n", proc.time()[["elapsed"]]-task_init))

if( !interactive() ) q("no", status=0)