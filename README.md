# ParadiseOS
This is a sort of "file system" implementation of the ParadiseOS project, albeit heavily modified. I wont be
giving the full explanation about the the stuff but i will link in the place where you can understand the concept, and
tell where my controls diverge from theirs. In resume, it seeks to simulate a sort of "virtual reality" where the user 
is a force and can manipulate objects around him by choosing to "become" one of those objects.

LINK: https://wiki.xxiivv.com/site/paradise.html

## Quick Check
a) "note", "look", "pass" and "learn" were removed.  
b) "program" was renamed to "learn".  
c) "script" and "follow" primitives were added.  
d) "create create" is allowed.  
e) "transform" destroy the current vessel.  
f) "transform" can be called on the habited vessel, in this case, the current vessel will be reseted.  
g) "script" 

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

This is what differs it from a directory/file behavior, which is cool, and make it feels i am play videogame.

## Proposed New Nomenclature
Name: FactoryOS  
User: Fabricator/Bot    
Primitives: buid - equip - access  

## Observations
1. This system does not only define input, but output behavior. The user must only see and interact(seeing is a form of interaction)
with vessels in the same hierarchy.  
