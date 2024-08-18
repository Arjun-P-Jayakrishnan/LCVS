#	LCVS

### INTRODUCTION

**LCVS** or Locally Controlled Version System is a version control system only to be used locally. The idea here is to make an application that has a bunch of features that let you create a versioning system but locally. This can be done using existing systems but i wanted to do in go as  a method to understand key core aspects in go.

### Concept

The **LCVS** has a similar approach to that of git but uses a hybrid approach. Instead of delta as in mercurial or as checkpointing in other versioning systems we intend to use a hybrid of both to do the work.
The system has almost the same approaches as git but allows user to store their version of product with description.


###	Dev Dependencies


####	Database
For the database we have chosen postgres and some related packages to handle our migrations and queries.

Here we have kept the folder structure to have a sql->schema-> files.sql format each with a up and down migrations.

##### SQL - Postgres
Follow instructions online .

##### SQL Migrations  - goose

To install go to 

    https://github.com/pressly/goose

To initiate a migration you must run this command

    goose postgres postgres://<postgres-user-name>:<password>@localhost:<port>/<project eg rssagg> <up or down>

Note:
the goose command may not work in windows so highly encourage you to download other or use a binary file and just paste the path to .exe file where 'goose' is in the command and keep the rest.

##### SQL Queries - sqlc

To install go to 

    https://docs.sqlc.dev/en/latest/overview

To run query

    sqlc

Note:
The sqlc command may not work in windows so highlyy encourage you to use other or download the binary .exe file and just paste the path to .exe file where 'sqlc' is in the command and keep the rest.
