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
the same time and space.

Bellow is show, that this new way better represents the logic design in place.

create garde

garden
tower
move tower garden
move red knight

The old way is not by itself inferior, just different. In my case i value logic uniformity and singularity, for personal taste and because
by following the a singular logic design to the end seems, at least to me, to... cant say into words.

This is reinforced by the fact that the old transform cannot allow the force to become a vessel
that already exist which for me, implied two logic designs in place 

This is what differs it from a directory/file behavior, which is cool, and make it feels i am play videogame.
// remember, this system does not only define input, but output behavior. the user must only see and interact(seeing is a form of interaction)
with vessels in the same hierarchy

## Proposed New Nomenclature
Name: FactoryOS  
User: Fabricator/Robot
Primitives: buid - equip - access  
