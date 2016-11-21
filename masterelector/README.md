This master election has one big difference from regular master election: it
prefers having a master over consistency. It is designed that every node has a
master, and if it doesn't it becomes a master itself. In case of netsplits this
will result in multple masters.

The rationale behind this is that we preferred multiple masters above no master.
Better to get a notification twice than not to get it at all.

The implementation will probably be that every node generates a random number
and the node with the highest number you're connected to is your master. This
might be yourself. In case of a netsplit, and you're not connected to any other
node, you will have the highest number you know of and you'll be the master.
