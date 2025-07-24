# ParadiseOS
This is a sort of "file system" implementation of the ParadiseOS project, albeit heavily modified. I wont be
giving the full explanation about the stuff but i will link in the place where you can understand the concept, and
tell where my controls diverge from theirs. In resume, it seeks to simulate a sort of "virtual reality" where the user 
is a force and can manipulate objects around him by choosing to "become" one of those objects.

Right now only the "follow" primitive is not totally implemented, i could not get a interpreter to work inside the code. So i will
be focusing on building one for the time being. For now it only prints the saved script.

LINK: https://wiki.xxiivv.com/site/paradise.html

## Quick Check
a) "note", "look", "pass" and "learn" were removed;  
b) "program" was renamed to "learn";  
c) "script" and "follow" primitives were added;  
d) "create create" is allowed;  
e) "transform" destroys the current vessel;  
f) "transform" can be called on the possessed vessel, in this case, the possessed vessel will be reseted;  
g) "warp" is warping the vessels that learned this skill, not the user;  
f) "script" saves a program written in Go;  
h) "follow" executes a given script;  
i) "script make garden : use = true" is the proper formatting, everything after ":" is program code. "garden:" is valid.

## "transform"
The reasons for the change in the transform primitive were two fold. First, the paradise logic design states that
each vessel is a entity that exist independently from one another and also from the "force". This is evident by the 
fact that such "force" can "become" a vessel, and afterwards, cease to be it without the annihilation of the vessel. 
As such, "becoming" here is analogous with wearing a vessel like cloth or "possessing" it, in exorcist terms. In fact 
not only it is implied that the vessel never ceases to be in the process of something becoming it, one cannot give 
the same name to two different vessels, even if they are inside different vessels, making so that the name is a 
representation of the essence of a object, it's individuality.  

As such, the old transform, if to follow the logic design in place, should had involved the destruction of the current 
possessed vessel and the creation of the new vessel already in a possessed state, in fact, the destruction produced 
by the transformation is fruit of the existential colision between two separated vessels that cannot coexist at the 
same time and space. The new transform as such is simple a sequence of anihilation -> creation. Because of this simple
definition i allowed one to transform into a vessel already in a possessed state, this vessel will be destroyed and 
then be called again into existence, fruit of the conflict between its existance and non-existance that cannot be at
the same time and space. Because of this I also made possible for one to transform into a already existent vessel.  

This new transform does three things: allows me to destroy vessels in a way i find elegant and without needing to create
a "destroy" primitive, allows me to reset a current possessed vessel and prevents a sequence of destructions that would 
lead to the force to be outside vessels.

The old way is not by itself inferior, just different. In my case i value logic uniformity and singularity, for personal 
taste and because by following a singular logic design to the end seems less confusing. For example, the old transform 
worked like a "rename" command of a file like system, where everything is a file, which has a name and so on. I want to
do away with such system.  

This is what differs it from a directory/file behavior, which is cool, and make it feels like i am playing a videogame.

## "warp"
Same reason as "transform". The logic design in place was that a force/user could teach vessels the primitives he knows. All primitives pointed to this except "warp". Aside from having the same name those primitives were not the same. I believe it was fruit from a desire to use it as a file like system shortcut, which mixed two logic designs and would blind me to the true way to use this vessels system. I am mostly certain that there are other ways of implementing a "portal". Or maybe not the old way itself, but the reasoning behind it, dont know how to explain this properly...  

## Proposed New Nomenclature
Name: FactoryOS  
User: Fabricator/Bot    
Primitives: buid - equip - access - leave - take - drop - allocate...  

## Observations
1. This system does not only define input, but output behavior. The user must only see and interact(seeing is a form of interaction)
with vessels in the same hierarchy.  
2. No form of visualizing the created scripts is to be programmed. I should find a way of registering this inside the "vessel" system, not outside.
3. Everything is supposed to be achieved thought the primitives of the paradise system, from coloring the monitor to producing sounds. Although possible to alter the states of input and output devices from the script, the idea is to program anything through these primitives. The script language is only to automate the execution of the primitives. I suspect this might be possible somehow...
4.  The idea of writing primitives as "create garden create tower move tower garden enter garden enter tower" in one go is rejected. You cant move something it does not exist yet, neither enter something outside your reach, like entering a bedroom before entering it's house. This system must simulate the interactive physical capabilities of a common person, which includes it's limitations. Aside from the appearances, the primitives represent the actions possible by the human being and not textual instructions. The primitives must be manual as the hands of a person.  
