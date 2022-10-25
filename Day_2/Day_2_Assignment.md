# Day 2 Assignment
Create GitHub repository using git for AquaDeveloper Batch 2 Wave 2 assignments

## Initialization
First, create directory in local machine by mkdir.

```bash
mkdir Academy
```

Since this Day_2_Assignment.md is commited in Day_2.

```bash
cd Academy
mkdir Day_2
```
In directory ../Academy we initialize git by git init.

```bash
git init
```
The terminal will show the following

```bash
PS C:\Users\Arymurti\Documents\eFishery\Academy> git init
Initialized empty Git repository in C:/Users/Arymurti/Documents/eFishery/Academy/.git/
```

As a result ../Academy folder will have a hidden folder called .git as seen below

![alt text](screenshot/git_init.png "git init succesful")

## Add & Commit
Whenever there are changes in local machine we need to use

```bash
git add .
git commit -m "Example Commit"
```
For example let us create a file called "test_notepad.txt" in ../Academy/Day_2 

By using **git add .** All changes in local directory will be added to local branch 

Next with **git commit -m "Example Commit"** The changes

```bash
PS C:\Users\Arymurti\Documents\eFishery\Academy> git add .
PS C:\Users\Arymurti\Documents\eFishery\Academy> git commit -m "Example Commit"
[master c6d975f] Example Commit
 2 files changed, 12 insertions(+), 2 deletions(-)
 create mode 100644 Day_2/test_notepad.txt   
```
We could see that several changes has been saved, including the creation of "test_notepad.txt" in folder Day_2.

## Connect to GitHub Repository

Before we push our changes to a GitHub repository. We must first create the repository. In this case I've created https://github.com/SantoAry/AquaDeveloper_Academy.

To connect to my repository, I use

```bash
git remote add origin https://github.com/SantoAry/AquaDeveloper_Academy.git  
```

## Pushing to Github Repository

By all the previous step, we can finally push our local files into our GitHub repository by

```bash
git push origin master
```
The result will be shown in the following in terminal