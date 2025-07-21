# ParadiseOS
This is a sort of "file system" implementation of the ParadiseOS project, albeit heavily modified. I wont be
giving the full explanation about the the stuff but i will link in the place where you can understand the concept, and
tell where my controls diverge from theirs.

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

The reasons for the change in the transform primitive were two fold. First, the paradise logic design state that
each vessel is a entity that exist independely from one another and also from the "force". This is evident by the fact
that the force can "become" a vessel, and afterwards, cease to be it without the anihilation of the vessel. As such, becoming
here is analogous with wearing a vessel like cloth or "possesing" it, in exorcist terms. In fact not only it is implied that
the vessel never ceases to be in the process of something becoming it, one cannot give the same name to two different vessels, even
if they are inside different vessels, making so that the name is a representation of the essence of a object, it's individuality. 
As such, the old transform, if to follow the logic design in place, should involve the destruction of the current posessed vessel and the creation
of the new vessel already in a possessed state. Bellow is show that this new logic design can represent such a fact.

example showing

garden
tower
move tower garden
move red knight

The old way is not by itself inferior, just different. In my case i value logic uniformity and singularity, for personal taste and because
by following the a singular logic design to the end seems, at least to me, to... cant say into words.


This is what differs it from a directory/file behavior


## Proposed New Nomenclature
Name: FactoryOS  
User: Fabricator/Robot
Primitives: buid - equip - access  
