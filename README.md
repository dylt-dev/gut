# gut news!

Git Internals and Das Ultimate Keyboard had a baby, and that baby is `gut`! VERY `gut`!

Git Internals does an incredible job of walking you through the 'plumbing' of git -- the low-level commands that do the dirty
work and allow the 'porcelain' -- the higher level commands -- to shine. These low-level commands make up the guts of git, and are 
the unsung heroes that make `commit`, `merge`, even `log` possible.

`gut` takes it up a notch. `gut` is a CLI replacement for `git`, that _only_ implements the low-level commands. It makes only the
guts of git available, forcing you to learn them like you've never learned them before.

Low-level commands are passed along to git (installation of the git CLI is required ... for now).
Other git commands will politely end with a reference to Git Internals.

Why would you do this? It's fun! In the words of the Git Internals author(s) ... 

"The content-addressable filesystem layer is amazingly cool"

Why would they lie?

More seriously, there's a good chance that `git` is extremely important to you or your team. That problems with git -- a bad merge, a corrupt repo or working copy -- are the most feared and crushing problems you might experience. And that despite that, no one really understands how git really works. At all.

You can fix that :)

### gut started!

```
% echo "gut vibes" | /tmp/gut hash-object -w --stdin

Recognized command: hash-object ([hash-object -w --stdin])
We gut this!

3816ef9b0d1888cfa5fc7937f2e174a4edf3c75d
~/dev/gut
% /tmp/gut cat-file -p 3816ef9b0d1888cfa5fc7937f2e174a4edf3c75d

Recognized command: cat-file ([cat-file -p 3816ef9b0d1888cfa5fc7937f2e174a4edf3c75d])
We gut this!

gut vibes
~/dev/gut
% /tmp/gut clone https://github.com/dylt.dev/gut/

Gut news! You can use simple commands to do what you need, and become invincible!
Go to https://git-scm.com/book/en/v2/Git-Internals-Plumbing-and-Porcelain to learn more.

Have a gut time!

~/dev/gut
% /tmp/gut something-else-entirely

Oh my gutness ... we don't recognize that command! Sorry!
```
