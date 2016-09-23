# First Cipher:
# gluhtlishjrvbadvyyplkaohavbyjpwolypzavvdlhrvuuleatlzzhnlzdpajoavcpnlulyljpwolyrlfdvykpzaolopkkluzftivsvmklhaoputfmhcvypalovsilpuluk

# Second Cipher:
# vwduwljudeehghyhubwklqjlfrxogilqgsohdvhuhwxuqdqbeoxhsulqwviruydxowdqgdodupghvljqedvhgrqzklfkedqnbrxghflghrqldpvhwwlqjxsvdihkrxvhfr

irb(main):161:0> a.split('').map{|i| i.ord-7 < 97 ? 26 + i.ord-7 : i.ord-7  }.map{|i| i.chr }.join
=> "zenameblackoutworriedthatourcipheristooweakonnextmessageswitchtovigenerecipherkeywordisthehiddensymbolofdeathinmyfavoriteholbeinend"

irb(main):172:0> b.split('').map{|i| i.ord-n < 97 ? 26 + i.ord-n : i.ord-n  }.map{|i| i.chr }.join
=> "startigrabbedeverythingicouldfindpleasereturnanyblueprintsforvaultandalarmdesignbasedonwhichbankyoudecideoniamsettingupsafehouseco"

