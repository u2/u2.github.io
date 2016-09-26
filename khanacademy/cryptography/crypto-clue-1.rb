# First Cipher:
# gluhtlishjrvbadvyyplkaohavbyjpwolypzavvdlhrvuuleatlzzhnlzdpajoavcpnlulyljpwolyrlfdvykpzaolopkkluzftivsvmklhaoputfmhcvypalovsilpuluk

# Second Cipher:
# vwduwljudeehghyhubwklqjlfrxogilqgsohdvhuhwxuqdqbeoxhsulqwviruydxowdqgdodupghvljqedvhgrqzklfkedqnbrxghflghrqldpvhwwlqjxsvdihkrxvhfr

irb(main):161:0> a.split('').map{|i| i.ord-7 < 97 ? 26 + i.ord-7 : i.ord-7  }.map{|i| i.chr }.join
=> "zenameblackoutworriedthatourcipheristooweakonnextmessageswitchtovigenerecipherkeywordisthehiddensymbolofdeathinmyfavoriteholbeinend"

irb(main):172:0> b.split('').map{|i| i.ord-n < 97 ? 26 + i.ord-n : i.ord-n  }.map{|i| i.chr }.join
=> "startigrabbedeverythingicouldfindpleasereturnanyblueprintsforvaultandalarmdesignbasedonwhichbankyoudecideoniamsettingupsafehouseco"


ze name blackout worried that our cipher is too weak on next message switch to vigenere cipher keyword is the hidden symbol of death in my favorite holbein end

start i grabbed everything i could find please return any blueprints for vault and alarm design based on which bank you decide on i am setting up safe house co
