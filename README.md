# ParadiseOS
This is a sort of "file system" implementation of the ParadiseOS project, albeit heavily modified. I wont be
giving the full explanation about the the stuff but i will link in the place where you can understand the concept, and
tell where my controls diverge from theirs.

LINK: https://wiki.xxiivv.com/site/paradise.html

## QuickCheck
a) "note", "look", "pass" and "learn" were removed.  
b) "program" renamed to "learn".  
c) "script" and "follow" primitives added.  
d) Vessels can be named after primitives, that is, "create create".

## Primitive "transform"
This one was heavily modified. The idea is that "transform" does not rename a vessel, instead, the current vessel
will be destroyed along all it's contents and anything else related to it. In this case, the name of the vessel
is being used to simulate it's individuality, it's defining concept. The idea is that the "force" is changing the form
it's assuming. With this, it's possible to have a way to destroy vessels without needing a "destroy" primitive. Also it's
possible for a vessel to transform into itself again, where it will cease to be and soon after, become again. This will 
destroy everything associated with the vessel and bring it back in a clean state. Basically i am defining "transform" as
a sort of "cease to be", to "become again" or "becom anew". Hmm... i guess a better idea is that "transform" involves the 
force assuming a new form? Perharps it is a better way of seeing?
