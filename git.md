## GITHUB CHEATSHEET 

Note: **markdown** are for short notes , refer <!-- "comments" --> for quick elaboration 

## basic 
git version <!-- to check the version of current git -->

git help <command name>  -> Help command to open any manual of an command <!-- Like <rebase> <push> <pull> -->

git config --global user.name/user.email <!-- followed by your name or eåmail --> "your name/email"

git config --global --list -> Shows all global Git settings

## REPO BASICS 

**Initialize** : git init
**Init with name** :git init <folder name>
**clone** : git clone <url> <!-- for folder add name after url--> statement-4
git clone -b <branch> <url> <!-- branch copy  & -b for branch -->

## status , stages

**check status** : git status 
git status -sb <!-- provides a short summary of repo current state-->

<!-- there r 3 areas mainly and the second one is staging where the main work is sent after creation of file -->

git add git.md(filname)
**add all changes** : git add . <!-- add all changes -->

N <!-- SAVE THE CHANGES INTO GIT HISTORY -->

## ## BRANCHING - creation,swtiching renaming and deletion <!-- already seen in creating a pr and also used other places -->

git branch <!-- shows the list of branches -->

* create + switch together 
it checkout -b branchname <!-- eg- issue4-statemeent-4-branch-fix -->
git switch -c branchname <!-- creates a branch and switches to it immeditaely -->
<!-- NOTE-> main diff in checkout and switch is that checkout can be sometime dangerous as it cna mistake a branch name with a same file name if exisits so safer way is to use switch and restore used to make a file say statement.txt to its prev state if u are unsure about the current one  -->

* only switch a branch
git checkout main
git switch main

## renaming 
**current branch** : git branch -m newname <!-- needs to be at the current branch to change the name , m stands for rename or move -->
**Any branch** : git branch -m oldname newname  <!-- need not to be on the current branch to change the branch name -->

## deletion 
** safe delete** : git branch -d branchname <!-- safe deletion prefers merge competion else gives warning -->
** force delete** : git branch -D  branchname <!-- forcefull deletion doesn't care if merge req are completed -->

## PULL AND PUSH 

**Simple pull** : git pull <!-- it pastes the new pulled code to the exisitng code -->
**Specific pull** : git pull origin main <!-- pulls specifically from origin -->
**Rebase pull** : git pull --rebase <!-- SYSTEMATIC OPTION for pull , it pulls the file and add on your content instead of merging them is unsystematic manner so the history is not messed up instead it looks polish -->
**Global setting** : git config --global pull.rebase true <!-- stores in history so that git comand is always --rebase no need to type it again its stored in the memory -->
git pull --all <!-- pull from all remotes -->

**intital push** : git push
git push -u origin <branchname> <!-- "-u" creates a relationship between local and github branch and pushes using this next time simply use git push -->

## REMOTE 

**View:** : git remote -v <!-- views the current location of local folder in github , like checking if statement-4is linked to chahat_internship under chahxtj github user or not -->
**Connect:** : git remote add origin <url> <!-- connects a new folder say newfolder1 to my github url , used first time then normal remote view -->
**Update:** : git remote <set-url> origin <new-url> <!-- used when repo or acc name or address is changed hence creates a connection directly -->
**Disconnect:** : git remote remove origin <!--cuts connection used when u copy coe from someone else's n needs to break the connection -->

## FETCHING
**checking** : git fetch <!-- checks code/changes from a repo w/o merging unlike pull-->
**fetch from origin** : git fetch origin 

## merge

git merge <branch> <!-- merge branches with main , eg issue8-statement4-branch merges with main branch , in history it creates a merge knot-->
git rebase main <!-- cleans history so it seems as a clear path , so no merge knots in path clearly replaces whole code to main -->


## stash 
git stash <!-- removes the uncommitted changes and stores at safe place hence clean work area - Temporarily saves your uncommitted work. -->
git stash list <!-- gives the stashes list -->

git stash apply <!-- copy back the stash to main but leaves history in stash list -->

git stash pop <!-- same as apply just removes the history from list-->

git stash clear : empty vault 

## RESTORE , RESET & REVERT

git restore <file> <!-- restore to last changes, removes the new changes-->

git restore --staged <file> <!-- removes files from git add stage -->

git restore . : undo all <!-- restores all unsaved changes and brings the file to state before the unsaved changes-->

git reset --soft <commitid> <!-- find id from log ,undoes commit but keeps your code changes staged  -->
git reset <commitid>  <!-- undoes commit and unstages changes, but keeps code in files -->
git reset --hard <commitid><!-- DANGER: deletes everything! commit and code both gone -->

git revert <commitID> <!-- creates a NEW commit that does the opposite of the target commits, so u have a new commit that has all data bedfore prev reset commit for eg -->
# LOGS & diff 

git log --oneline <!--commit history in short one-line format -->
git show <commit> <!-- full details of a specific commit files changed n code diff -->
git blame <file> <!-- see's modified each line of a file and in which commit -->
git diff <!-- shows unstaged changes -->
git diff --staged <!-- shows staged changes compared to last commit -->
git diff main..feature-x <!-- compares differences between main and feature-x branches -->

## RESET / REVERT

git reset --soft HEAD~1 <!-- removes last commit but keeps changes staged -->
git reset --hard HEAD~1 <!-- removes last commit and deletes all its changes permanently -->
git revert HEAD <!-- undoes last commit safely by creating a new reverse commit -->
git commit --amend -m "new msg" <!-- edits the last commit message -->


## CLEAN

git clean -n <!-- shows which untracked files will be deleted -->
git clean -f <!-- force deletes untracked files -->
git clean -fd <!-- force deletes untracked files and directories -->

## TAGS
git tag v1.0.0 <!--creates a lightweight tag -->
git tag -a v1.0.0 -m "release" <!-- adds annotated tag with message-->
git push origin v1.0.0 <!-- pushes the tag to remote repository -->

## FILE MANAGEMENT

git rm <file> <!-- delete file & stage -->
git mv <old> <new> <!-- move or rename file & stage -->

## SHORTCUTS
git config --global alias.st "status -sb" <!-- creates shortcut 'git st' for short status view -->
git config --global alias.lg "log --oneline --decorate --graph --all" <!-- creates shortcut 'git lg' for graphical log view -->
git st <!-- shortcut command for git status -sb -->
git lg <!-- shortcut command for graphical log view -->


## PERSONAL SHORTCUTS

git config --global alias.co "checkout" <!-- used for checkout main branch -->
git config --global alias.br "branch" <!--  git branch -->
git config --global alias.ci "commit" <!-- shortcut for git commit -->
git config --global alias.last "log -1 HEAD" <!-- shows last commit details -->
git config --global alias.unstage "reset HEAD --" <!-- unstages a file -->
git config --global alias.undo "reset --soft HEAD~1" <!-- undo last commit but keep changes -->
git config --global alias.graph "log --oneline --graph --decorate --all" <!-- clean visual branch graph -->
git config --global alias.s "status" <!-- shorter status command -->
git config --global alias.a "add ." <!-- adds all changes quickly -->
git config --global alias.cm "commit -m" <!-- quick commit message: git cm "msg" -->
git config --global alias.amend "commit --amend --no-edit" <!-- amend last commit without changing message -->
git config --global alias.recent "branch --sort=-committerdate" <!-- shows recently updated branches -->
git config --global alias.save "!git add -A && git commit -m 'save'" <!-- quick save everything -->

