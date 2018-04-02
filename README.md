## DANGER: NOT EVEN REMOTELY CLOSE TO BEING READY TO USE.  (But you're welcome to have a look)

# Trebuchet
A tool to scaffold projects quickly.  Similiar to others, Ponzu, Buffalo, Fragmenta, etc.  Not intended to be a replacement to 
such projects, but something aimed at how I personally like to do things.  

# To Use It
This will be updated as the project moves further along.  Currently it's just a hodge-podge of scatter brained folders that might
be all tied together one day. Currently, the only thing `main.go` does right now is nothing.  It's set to execute whatever root 
command is also passed along with the program.

# Inside the app folder
  - app.go - This just grabs some initial configuration settings from `config.json` and starts a server based on gorilla mux.
  
  - configuration folder - Holds config.json and an example config file.  No where near complete, just building it out as I
    progress.  This will kind of serve as the `.env` some might be used to seeing in other toolsets.
    
  - datalayer - Empty, I'm assuming I will eventually use this for everything DB related.  I just needed to remind myself of
    the thing I still need to get to.

  - hash - contains `hmac.go` which has some stuff for encypting passwords

  - models - models will be stored here.
    
  - playground - where I go to screw things up.  It's really just a testing ground when I am working on stuff, which is why most
    things are commented out, almost completely.
    
# Inside the cmd folder
  - all of the files in here are just commands for the trebuchet command line tool. Each command will have it's own file..this
    is just following whatever they told me at spf/cobra.  I honestly have no idea what I'm doing.
    
# Inside the hash folder
  - Stuff to hash passwords etc...etc... nothing is really wired together yet, but one day
  
# Inside the trebuchet folder
  - You know, I don't really know what that folder is there.  I just always see people's stuff and it's always got some kind of
    doubled up folder structure.  More to come on this when I figure it out.
