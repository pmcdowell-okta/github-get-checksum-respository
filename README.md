# github-get-checksum-respository

##Get the Master Checksum of a Github Repository

I wrote this to I could easily check and see if a github repository has changed.

I'm going to add a for/loop, to wait for a change, if there is a change it will
exit. The idea is I can run this on a server that is remote, so if something
changes I can pull the changes and restart the service.

I'm sure there is an easier way to do this, but it was the best I could come up with
