package scriptures

import (
    "fmt"
    "time"
    "bcolors"
    "math/rand"
)

type Scriptures struct{}

func TheScriptures() {
    fmt.Printf(`%s
                     _,._
                 __.'   _)
                <_,)'.-"a\
                  /' (   \
      _.-----..,-'   ('"--^
     //              |
    (|   ';      ,   |
      \   ;.----/  ,/
       ) // /   | |\ \
       \ \\'\   | |/ /      %sJesus Christ%s
        \ \\ \  | |\/  %sThe Lamb that was slain%s
         '" '"  '""         %sfor our sins.%s`, bcolors.BLUE, bcolors.GREEN + bcolors.GREEN, 
bcolors.BLUE, bcolors.GREEN, bcolors.BLUE, bcolors.GREEN,
 bcolors.ENDC)
    fmt.Printf(bcolors.YELLOW + `
         __                 _____ _____     _     _
      __|  |___ ___ _ _ ___|     |  |  |___|_|___| |_
     |  |  | -_|_ -| | |_ -|   --|     |  _| |_ -|  _|
     |_____|___|___|___|___|_____|__|__|_| |_|___|_|
                         ¯\_(ツ)_/¯
` + bcolors.ENDC)
     fmt.Printf(bcolors.GREEN + `
             THE BIBLE VERSES TO GUIDE YOU HOME. ` + bcolors.ENDC)
    scriptures := bcolors.DARKCYAN + `
(In the beginning God created the heaven & the earth (Gen.1:1)
(& God saw the light, that it was good: & God divided the light from the darkness. (Gen.1:4)
(I will bless thee…  & thou shalt be a blessing (Gen.12:2)
(I am thy shield, & thy exceeding great reward (Gen.15:1)
(Fear ye not, st& still, & see the salvation of the LORD (Ex.14:13)
(I will make all My goodness pass before thee (Ex.33:19)
(The LORD God, merciful & gracious (Ex.34:6)
(I set my tabernacle among you (Lev.26:11)
(I will walk among you, & will be your God (Lev.26:12)
(The LORD is longsuffering, & of great mercy (Num.14:18)
(Thou shalt love the LORD thy God with all thine heart (Deut.6:5)
(Shall ye lay up these my words in your heart & in your soul (Deut.11:18)
(Thou shalt rejoice before the LORD thy God (Deut.12:18)
(Blessed shalt thou be in the city, & blessed shalt thou be in the field (Deut.28:3)
(Blessed shall be the fruit of thy body, & the fruit of thy ground (Deut.28:4)
(Blessed shall be thy basket & thy store (Deut.28:5)
(Blessed shalt thou be when thou comest in, & blessed shalt thou be when thou goest out (Deut.28:6)
(They shall come out against thee one way, & flee before thee seven ways (Deut.28:7)
(& the LORD shall make thee the head, & not the tail (Deut.28:13)
(Be strong & of a good courage, fear not, nor be afraid (Deut.31:6)
(I will not fail thee, nor forsake thee (Josh.1:5)
(Only be thou strong & very courageous (Josh.1:7)
(This Book of the Law shall not depart out of thy mouth (Josh.1:8)
(Have not I comm&ed thee? Be strong & of a good courage (Josh.1:9)
(Be not afraid, neither be thou dismayed: for the LORD thy God is with thee (Josh.1:9)
(Sanctify yourselves: for to morrow the LORD will do wonders among you (Josh.3:5)
(The LORD your God, He it is that fighteth for you (Josh.23:10)
(Nay; but we will serve the LORD! (Josh.24:21)
(I will never break my covenant with you (Judg.2:1)
(Them that love him be as the sun when he goeth forth in his might (Judg.5:31)
(Thy people shall be my people, & thy God my God (Ruth 1:16)
(He will keep the feet of His saints (1Sam.2:9)
(Only fear the LORD, & serve Him in truth with all your heart (1Sam.12:24)
(I will shew thee what thou shalt do (1Sam.16:3)
(Man looketh on the outward appearance, but the LORD looketh on the heart (1Sam.16:7)
(Yet he hath made with me an everlasting covenant (2Sam.23:5)
(Give therefore thy servant an underst&ing heart (1Kin.3:9)
(The LORD your God ye shall fear; & he shall deliver you (2Kin.17:39)
(Serve Him with a perfect heart & with a willing mind (1Chr.28:9)
(The LORD searcheth all hearts, & underst&eth all the imaginations of the thoughts (1Chr.28:9)
(If thou seek him, he will be found of thee (1Chr.28:9)
(The LORD is able to give thee much more than this (2Chr.25:9)
(Will not turn away His face from you, if ye return unto him (2Chr.30:9)
(He did it with all his heart, & prospered (2Chr.31:21)
(The hand of our God is upon all them for good that seek Him (Ezdra 8:22)
(The joy of the LORD is your strength (Neh.8:10)
(& who knoweth whether thou art come to the kingdom for such a time as this? (Esther.4:14)
(Till he fill thy mouth with laughing, & thy lips with rejoicing (Job 8:21)
(I know that my redeemer liveth (Job 19:25)
(Blessed is the man that walketh not in the counsel of the ungodly (Ps.1:1)
(The LORD knoweth the way of the righteous (Ps.1:6)
(Serve the LORD with fear, & rejoice with trembling (Ps.2:11)
(Blessed are all they that put their trust in Him (Ps.2:12)
(But thou, O LORD, art a shield for me, my glory (Ps.3:3)
(Lead me, O LORD, in thy righteousness (Ps.5:8)
(For thou, LORD, wilt bless the righteous (Ps.5:12)
(For the righteous God trieth the hearts (Ps.7:9)
(I will praise thee, O LORD, with my whole heart (Ps.9:1)
(For thou, LORD, hast not forsaken them that seek Thee (Ps.9:10)
(The LORD is King for ever & ever (Ps.10:16)
(LORD, thou hast heard the desire of the humble (Ps.10:17)
(The LORD is the portion of mine inheritance & of my cup (Ps.16:5)
(Hide me under the shadow of thy wings (Ps.17:8)
(I will love thee, O LORD, my strength (Ps.18:1)
(The LORD is my rock, & my fortress, & my deliverer (Ps.18:2)
(For thou wilt light my c&le (Ps.18:28)
(The LORD my God will enlighten my darkness (Ps.18:28)
(It is God that girdeth me with strength (Ps.18:32)
(The LORD liveth; & blessed be my rock (Ps.18:46)
(Let the God of my salvation be exalted (Ps.18:46)
(The heavens declare the glory of God; & the firmament sheweth His handywork (Ps.19:1)
(The law of the LORD is perfect, converting the soul (Ps.19:7)
(The testimony of the LORD is sure, making wise the simple (Ps.19:7)
(The statutes of the LORD are right, rejoicing the heart (Ps.19:8)
(The comm&ment of the LORD is pure, enlightening the eyes (Ps.19:8)
(The fear of the LORD is clean, enduring for ever (Ps.19:9)
(The judgments of the LORD are true & righteous altogether (Ps.19:9)
(Some trust in chariots, & some in horses: but we will remember the name of the LORD our God. (Ps.20:7)
(They shall praise the LORD that seek Him (Ps.22:26)
(All the ends of the world shall remember & turn unto the LORD (Ps.22:27)
(A seed shall serve Him (Ps.22:30)
(The LORD is my shepherd; I shall not want (Ps.23:1)
(The LORD strong & mighty, the LORD mighty in battle (Ps.24:8)
(O my God, I trust in Thee: let me not be ashamed (Ps.25:2)
(Shew me thy ways, O LORD; teach me Thy paths. (Ps.25:4)
(Good & upright is the LORD: therefore will He teach sinners in the way (Ps.25:8)
(The secret of the LORD is with them that fear Him (Ps.25:14)
(Mine eyes are ever toward the LORD; for He shall pluck my feet out of the net. (Ps.25:15)
(Teach me Thy way, O LORD (Ps.27:11)
(Wait on the LORD: be of good courage, & He shall strengthen thine heart (Ps.27:14)
(The LORD is my strength & my shield (Ps.28:7)
(My heart trusted in Him, & I am helped (Ps.28:7)
(The voice of the LORD is powerful; the voice of the LORD is full of majesty (Ps.29:4)
(The LORD will give strength unto His people (Ps.29:11)
(The LORD will bless his people with peace (Ps.29:11)
(In his favour is life (Ps.30:5)
(Be of good courage, & he shall strengthen your heart (Ps.31:24)
(Blessed is he whose transgression is forgiven (Ps.32:1)
(I will instruct thee & teach thee in the way which thou shalt go (Ps.32:8)
(He that trusteth in the LORD, mercy shall compass him about (Ps.32:10)
(Be glad in the LORD, & rejoice (Ps.32:11)
(For the word of the LORD is right; & all His works are done in truth (Ps.33:4)
(The earth is full of the goodness of the LORD (Ps.33:5)
(Our soul waiteth for the LORD: He is our help & our shield. (Ps.33:20)
(They looked unto Him, & were lightened (Ps.34:5)
(O taste & see that the LORD is good (Ps.34:8)
(They that seek the LORD shall not want any good thing (Ps.34:10)
(The LORD is nigh unto them that are of a broken heart (Ps.34:18)
(& my soul shall be joyful in the LORD: it shall rejoice in His salvation. (Ps.35:9)
(Trust in the LORD, & do good (Ps.37:3)
(Dwell in the l&, & verily thou shalt be fed (Ps.37:3)
(Delight thyself also in the LORD; & He shall give thee the desires of thine heart (Ps.37:4)
(Commit thy way unto the LORD; trust also in Him; & He shall bring it to pass. (Ps.37:5)
(Rest in the LORD, & wait patiently for Him (Ps.37:7)
(The meek shall inherit the earth; & shall delight themselves in the abundance of peace (Ps.37:11)
(Wait on the LORD, & keep His way (Ps.37:34)
(The salvation of the righteous is of the LORD: He is their strength in the time of trouble. (Ps.37:39)
(Lord, all my desire is before thee (Ps.38:9)
(Blessed is that man that maketh the LORD his trust (Ps.40:4)
(I delight to do Thy will, O my God (Ps.40:8)
(Be still, & know that I am God (Ps.46:10)
(O clap your hands, all ye people; shout unto God with the voice of triumph! (Ps.47:1)
(The redemption of their soul is precious (Ps.49:9)
(Our God shall come, & shall not keep silence (Ps.50:3)
(Offer unto God thanksgiving; & pay thy vows unto the Most High (Ps.50:14)
(Call upon me in the day of trouble: I will deliver thee (Ps.50:15)
(Whoso offereth praise glorifieth Me (Ps.50:23)
(The fool hath said in his heart, "There is no God" (Ps.53:1)
(Behold, God is mine helper: the Lord is with them that uphold my soul (Ps.54:4)
(Cast thy burden upon the LORD, & He shall sustain thee (Ps.55:22)
(He shall never suffer the righteous to be moved (Ps.55:22)
(Verily there is a reward for the righteous (Ps.58:11)
(God is my defence, & the God of my mercy (Ps.59:17)
(Truly my soul waiteth upon God: from Him cometh my salvation (Ps.62:1)
(God is a refuge for us (Ps.62:8)
(God hath spoken once; twice have I heard this; that power belongeth unto God. (Ps.62:11)
(The righteous shall be glad in the LORD, & shall trust in Him (Ps.64:10)
(Thy God hath comm&ed thy strength (Ps.68:28)
(But it is good for me to draw near to God (Ps.73:28)
(Open thy mouth wide, & I will fill it (Ps.81:10)
(The LORD God is a sun & shield (Ps.84:11)
(The LORD will give grace & glory (Ps.84:11)
(The LORD shall give that which is good; & our l& shall yield her increase. (Ps.85:12)
(For thou, Lord, art good, & ready to forgive; & plenteous in mercy (Ps.86:5)
(He that dwelleth in the secret place of the Most High shall abide under the shadow of the Almighty (Ps.91:1)
(Thou shalt not be afraid for the terror by night; nor for the arrow that flieth by day (Ps.91:5)
(Lest thou dash thy foot against a stone (Ps.91:12)
(The righteous shall flourish like the palm tree (Ps.92:12)
(The LORD reigneth, He is clothed with majesty (Ps.93:1)
(The LORD knoweth the thoughts of man, that they are vanity (Ps.94:11)
(The LORD will not cast off His people (Ps.94:14)
(Give unto the LORD the glory due unto His name (Ps.96:8)
(Worship the LORD in the beauty of holiness (Ps.96:9)
(The LORD reigneth: the world also shall be established that it shall not be moved (Ps.96:10)
(Light is sown for the righteous, & gladness for the upright in heart (Ps.97:11)
(The LORD hath made known His salvation (Ps.98:2)
(Serve the LORD with gladness (Ps.100:2)
(I will set no wicked thing before mine eyes (Ps.101:3)
(They shall perish, but thou shalt endure (Ps.102:26)
(Bless the LORD, O my soul, & forget not all His benefits (Ps.103:2)
(Forget not all His benefits (Ps.103:2)
(Who forgiveth all thine iniquities; who healeth all thy diseases (Ps.103:3)
(As the heaven is high above the earth, so great is His mercy (Ps.103:11)
(As far as the east is from the west, so far hath He removed our transgressions from us (Ps.103:12)
(I will sing unto the LORD as long as I live (Ps.104:33)
(Let the heart of them rejoice that seek the LORD! (Ps.105:3)
(Seek the LORD, & His strength: seek his face evermore! (Ps.105:4)
(O give thanks unto the LORD; for He is good: for His mercy endureth for ever. (Ps.106:1)
(He sent his word, & healed them (Ps.107:20)
(Through God we shall do valiantly (Ps.108:13)
(I will greatly praise the LORD with my mouth (Ps.109:30)
(He raiseth up the poor out of the dust, & lifteth the needy out of the dunghill (Ps.113:7)
(Ye that fear the LORD, trust in the LORD: He is their help & their shield (Ps.115:11)
(He will bless them that fear the LORD (Ps.115:13)
(I will take the cup of salvation, & call upon the name of the LORD (Ps.116:13)
(His merciful kindness is great toward us (Ps.117:2)
(It is better to trust in the LORD than to put confidence in man (Ps.118:8)
(The LORD is my strength & song, & is become my salvation (Ps.118:14)
(The right hand of the LORD doeth valiantly (Ps.118:16)
(This is the day which the LORD hath made; we will rejoice & be glad in it (Ps.118:24)
(Thy word is a lamp unto my feet, & a light unto my path (Ps.119:105)
(Great peace have they which love thy law (Ps.119:165)
(My help cometh from the LORD, which made heaven & earth (Ps.121:2)
(He will not suffer thy foot to be moved (Ps.121:3)
(The LORD is thy keeper (Ps.121:5)
(I was glad when they said unto me, "Let us go into the house of the LORD" (Ps.122:1)
(Our soul is escaped as a bird out of the snare of the fowlers (Ps.124:7)
(The snare is broken, & we are escaped (Ps.124:7)
(Our help is in the name of the LORD (Ps.124:8)
(They that trust in the LORD shall be as Mount Zion (Ps.125:1)
(They that sow in tears shall reap in joy (Ps.126:5)
(Lift up your hands… & bless the LORD (Ps.134:2)
(The LORD will perfect that which concerneth me (Ps.138:8)
(The LORD upholdeth all that fall (Ps.145:14)
(The LORD is righteous in all His ways (Ps.145:17)
(The LORD is nigh unto all them that call upon Him (Ps.145:18)
(He healeth the broken in heart, & bindeth up their wounds (Ps.147:3)
(The LORD lifteth up the meek (Ps.147:6)
(Let the children of Zion be joyful in their King (Ps.149:2)
(For the LORD taketh pleasure in His people (Ps.149:4)
(Let every thing that hath breath praise the LORD (Ps.150:6)
(The fear of the LORD is the beginning of knowledge (Prov.1:7)
(My son, if sinners entice thee, consent thou not (Prov.1:10)
(I will pour out my spirit unto you, I will make known my words unto you (Prov.1:23)
(For the LORD giveth wisdom: out of His mouth cometh knowledge & underst&ing (Prov.2:6)
(He layeth up sound wisdom for the righteous (Prov.2:7)
(He keepeth the paths of judgment (Prov.2:8)
(Walk in the way of good men, & keep the paths of the righteous (Prov.2:20)
(He blesseth the habitation of the just (Prov.3:33)
(I give you good doctrine, forsake ye not my law (Prov.4:2)
(Get wisdom: & with all thy getting get underst&ing. (Prov.4:7)
(I have taught thee in the way of wisdom; I have led thee in right paths. (Prov.4:11)
(When thou goest, thy steps shall not be straitened (Prov.4:12)
(Take fast hold of instruction (Prov.4:13)
(The path of the just is as the shining light (Prov.4:18)
(Attend to my words; incline thine ear unto my sayings (Prov.4:20)
(My words… are life unto those that find them, & health to all their flesh (Prov.4:22)
(Keep thy heart with all diligence (Prov.4:23)
(Ponder the path of thy feet, & let all thy ways be established (Prov.4:26)
(Drink waters out of thine own cistern, & running waters out of thine own well (Prov.5:15)
(For the ways of man are before the eyes of the LORD (Prov.5:21)
(For the comm&ment is a lamp; & the law is light (Prov.6:23)
(My son, keep my words (Prov.7:1)
(Keep my commandments, & live; & my law as the apple of thine eye (Prov.7:2)
(Wisdom is better than rubies (Prov.8:11)
(Those that seek me early shall find me (Prov.8:17)
(Blessed are they that keep my ways (Prov.8:32)
(The LORD will not suffer the soul of the righteous to famish (Prov.10:3)
(The hand of the diligent maketh rich (Prov.10:4)
(The blessing of the LORD, it maketh rich, & He addeth no sorrow with it (Prov.10:22)
(The desire of the righteous shall be granted (Prov.10:24)
(The hope of the righteous shall be gladness (Prov.10:28)
(By the blessing of the upright the city is exalted (Prov.11:11)
(He that watereth shall be watered also himself (Prov.11:25)
(He that tilleth his l& shall be satisfied with bread (Prov.12:11)
(He that keepeth his mouth keepeth his life (Prov.13:3)
(The light of the righteous rejoiceth (Prov.13:9)
(He that walketh with wise men shall be wise (Prov.13:20)
(The wisdom of the prudent is to underst& his way (Prov.14:8)
(The tabernacle of the upright shall flourish (Prov.14:11)
(A true witness delivereth souls (Prov.14:25)
(The fear of the LORD is a fountain of life (Prov.14:27)
(A sound heart is the life of the flesh (Prov.14:30)
(A soft answer turneth away wrath (Prov.15:1)
(The eyes of the LORD are in every place (Prov.15:3)
(He that refuseth instruction despiseth his own soul (Prov.15:32)
(The fear of the LORD is the instruction of wisdom (Prov.15:33)
(Before honour is humility (Prov.15:33)
(Commit thy works unto the LORD, & thy thoughts shall be established (Prov.16:3)
(He that ruleth his spirit than he that taketh a city (Prov.16:32)
(The furnace for gold: but the LORD trieth the hearts (Prov.17:3)
(A merry heart doeth good like a medicine (Prov.17:22)
(A broken spirit drieth the bones (Prov.17:22)
(He that hath knowledge spareth his words (Prov.17:27)
(The name of the LORD is a strong tower (Prov.18:10)
(Death & life are in the power of the tongue (Prov.18:21)
(A man that hath friends must shew himself friendly (Prov.18:24)
(The discretion of a man deferreth his anger (Prov.19:11)
(He that keepeth the comm&ment keepeth his own soul (Prov.19:16)
(He that hath pity upon the poor lendeth unto the LORD (Prov.19:17)
(It is an honour for a man to cease from strife (Prov.20:3)
(Wait on the LORD, & he shall save thee (Prov.20:22)
(The LORD pondereth the hearts (Prov.21:2)
(A good name is rather to be chosen than great riches (Prov.22:1)
(He that hath a bountiful eye shall be blessed (Prov.22:9)
(Drowsiness shall clothe a man with rags (Prov.23:21)
(Through wisdom is an house builded; & by underst&ing it is established (Prov.24:3)
(By wise counsel thou shalt make thy war (Prov.24:6)
(In multitude of counsellors there is safety (Prov.24:6)
(For a just man falleth seven times, & riseth up again (Prov.24:16)
(A soft tongue breaketh the bone (Prov.25:15)
(The righteous are bold as a lion (Prov.28:1)
(When righteous men do rejoice, there is great glory (Prov.28:12)
(He that tilleth his l& shall have plenty of bread (Prov.28:19)
(The righteous considereth the cause of the poor (Prov.29:7)
(Honour shall uphold the humble in spirit (Prov.29:23)
(To every thing there is a season, & a time to every purpose under the heaven (Eccl.3:1)
(A time to cast away stones, & a time to gather stones together (Eccl.3:5)
(A time to keep silence, & a time to speak (Eccl.3:7)
(Whatsoever God doeth, it shall be for ever (Eccl.3:14)
(A threefold cord is not quickly broken (Eccl.4:12)
(Keep thy foot when thou goest to the house of God, & be more ready to hear (Eccl.5:1)
(Seeing there be many things that increase vanity (Eccl.6:11)
(Anger resteth in the bosom of fools (Eccl.7:9)
(God hath made man upright; but they have sought out many inventions (Eccl.7:29)
(A man's wisdom maketh his face to shine (Eccl.8:1)
(A wise man's heart discerneth both time & judgment (Eccl.8:5)
(It shall be well with them that fear God (Eccl.8:12)
(Let thy head lack no ointment (Eccl.9:8)
(Whatsoever thy hand findeth to do, do it with thy might (Eccl.9:10)
(Wisdom is better than weapons of war (Eccl.9:18)
(He that diggeth a pit shall fall into it (Eccl.10:8)
(Whoso breaketh an hedge, a serpent shall bite him (Eccl.10:8)
(Cast thy bread upon the waters (Eccl.11:1)
(He that observeth the wind shall not sow (Eccl.11:4)
(Remember now thy Creator in the days of thy youth (Eccl.12:1)
(Vanity of vanities, saith the Preacher; all is vanity (Eccl.12:8)
(Fear God, & keep his commandments: for this is the whole duty of man (Eccl.12:13)
(Take us the foxes, the little foxes, that spoil the vines (Song 2:15)
(Many waters cannot quench love (Song 8:7)
(Learn to do well (Is.1:17)
(Though your sins be as scarlet, they shall be as white as snow (Is.1:18)
(He will teach us of his ways, & we will walk in His paths (Is.2:3)
(Whom shall I send, & who will go for Us? (Is.6:8)
(Then said I, Here am I; send me. (Is.6:8)
(The LORD JEHOVAH is my strength & my song (Is.12:2)
(In the LORD JEHOVAH is everlasting strength (Is.26:4)
(My people shall dwell in a peaceable habitation (Is.32:18)
(In the wilderness shall waters break out, & streams in the desert (Is.35:6)
(He shall feed his flock like a shepherd (Is.40:11)
(He giveth power to the faint; & to them that have no might He increaseth strength (Is.40:29)
(They that wait upon the LORD shall renew their strength (Is.40:31)
(They shall run, & not be weary; & they shall walk, & not faint (Is.40:31)
(I the LORD thy God will hold thy right hand (Is.41:13)
(Fear not; I will help thee (Is.41:13)
(I will make the wilderness a pool of water (Is.41:18)
(I will make darkness light before them (Is.42:16)
(I will even make a way in the wilderness, & rivers in the desert (Is.43:19)
(I, even I, am he that blotteth out thy transgressions  (Is.43:25)
(Return unto Me; for I have redeemed thee (Is.44:22)
(The word is gone out of my mouth in righteousness, & shall not return (Is.45:23)
(I am the LORD thy God which teacheth thee to profit (Is.48:17)
(I have put my words in thy mouth (Is.51:16)
(I have covered thee in the shadow of Mine hand (Is.51:16)
(He is despised & rejected of men; a Man of sorrows, & acquainted with grief (Is.53:3)
(Surely he hath borne our griefs, & carried our sorrows (Is.53:4)
(He was wounded for our transgressions, He was bruised for our iniquities (Is.53:5)
(The chastisement of our peace was upon Him (Is.53:5)
(With his stripes we are healed (Is.53:5)
(The LORD hath laid on him the iniquity of us all (Is.53:6)
(He shall see of the travail of His soul, & shall be satisfied (Is.53:11)
(He bare the sin of many, & made intercession for the transgressors (Is.53:12)
(Enlarge the place of thy tent (Is.54:2)
(My kindness shall not depart from thee (Is.54:10)
(No weapon that is formed against thee shall prosper (Is.54:17)
(Thou shalt be like a watered garden (Is.58:11)
(The LORD shall be thine everlasting light (Is.60:20)
(I all their affliction he was afflicted (Is.63:9)
(Ask for the old paths, where is the good way, & walk therein (Jer.6:16)
(I am with thee to save thee & to deliver thee (Jer.15:20)
(I the LORD search the heart, I try the reins (Jer.17:10)
(Ye shall seek Me, & find Me, when ye shall search for Me with all your heart (Jer.29:13)
(I will forgive their iniquity, & their sin I will remember no more (Jer.31:34)
(I am the LORD, the God of all flesh: is there any thing too hard for Me? (Jer.32:27)
(I will rejoice over them to do them good (Jer.32:41)
(Call unto Me, & I will answer thee (Jer.33:3)
(I will shew thee great & mighty things, which thou knowest not (Jer.33:3)
(Obey, I beseech thee, the voice of the LORD (Jer.38:20)
(Come, & let us join ourselves to the LORD (Jer.50:5)
(The LORD is good unto them that wait for him (Lam.3:25)
(Let us search & try our ways, & turn again to the LORD (Lam.3:40)
(Let us lift up our heart with our hands unto God in the heavens (Lam.3:41)
(Cast away from you all your transgressions, whereby ye have transgressed (Ezek.18:31)
(I will accept you with your sweet savour (Ezek.20:41)
(I will cause the shower to come down in his season (Ezek.34:26)
(There shall be showers of blessing (Ezek.34:26)
(A new heart also will I give you, & a new spirit will I put within you (Ezek.36:26)
(Our God whom we serve is able to deliver us (Dan.3:17)
(For I desired mercy, & not sacrifice (Hos.6:6)
(Turn ye even to me with all your heart (Joel 2:12)
(Be glad & rejoice in the LORD your God (Joel 2:23)
(I will pour out My Spirit upon all flesh (Joel 2:28)
(Whosoever shall call on the name of the LORD shall be delivered (Joel 2:32)
(The LORD will be the hope of his people (Joel 3:16)
(Seek ye me, & ye shall live (Amos 5:4)
(The day of the LORD is near upon all the heathen (Obad.1:15)
(Salvation is of the LORD (Jon.2:9)
(We will walk in the name of the LORD our God (Mic.4:5)
(I will wait for the God of my salvation (Mic.7:7)
(The LORD knoweth them that trust in him (Nah.1:7)
(The just shall live by his faith (Hab.2:4)
(O LORD, revive thy work in the midst of the years (Hab.3:2)
(The LORD thy God in the midst of thee is mighty; He will save (Zeph.3:17)
(Thus saith the LORD of hosts; Consider your ways. (Hag.1:7)
(Be strong… for I am with you, saith the LORD (Hag.2:4)
(My spirit remaineth among you: fear ye not! (Hag.2:5)
(Fear not, but let your hands be strong (Zech.8:13)
(For I am the LORD, I change not (Mal.3:6)
(Bring ye all the tithes into the storehouse (Mal.3:10)
(& all nations shall call you blessed (Mal.3:12)
(Unto you that fear My name shall the Sun of Righteousness arise (Mal.4:2)
(Prepare ye the way of the Lord, make His paths straight (Matt.3:3)
(Therefore fruits meet for repentance (Matt.3:8)
(He shall baptize you with the Holy Ghost, & with fire (Matt.3:11)
(He will throughly purge His floor, & gather His wheat into the garner (Matt.3:12)
(Man shall not live by bread alone (Matt.4:4)
(Thou shalt worship the LORD thy God (Matt.4:10)
(Repent: for the kingdom of heaven is at hand (Matt.4:17)
(Follow Me, & I will make you fishers of men (Matt.4:19)
(Blessed are the poor in spirit: for theirs is the kingdom of heaven (Matt.5:3)
(Blessed are they that mourn: for they shall be comforted (Matt.5:4)
(Blessed are the meek: for they shall inherit the earth (Matt.5:5)
(Blessed are the merciful: for they shall obtain mercy (Matt.5:7)
(Blessed are the pure in heart: for they shall see God. (Matt.5:8)
(Blessed are the peacemakers: for they shall be called the children of God (Matt.5:9)
(Ye are the light of the world. A city that is set on an hill cannot be hid. (Matt.5:14)
(Let your light so shine before men (Matt.5:16)
(Be ye therefore perfect, even as your Father which is in heaven is perfect. (Matt.5:48)
(Your Father knoweth what things ye have need of (Matt.6:8)
(But lay up for yourselves treasures in heaven (Matt.6:20)
(For where your treasure is, there will your heart be also (Matt.6:21)
(But seek ye first the kingdom of God, & his righteousness (Matt.6:33)
(Ask, & it shall be given you (Matt.7:7)
(Seek, & ye shall find (Matt.7:7)
(Knock, & it shall be opened unto you (Matt.7:7)
(Your Father which is in heaven give good things to them that ask him (Matt.7:11)
(Narrow is the way, which leadeth unto life (Matt.7:14)
(Be of good cheer; thy sins be forgiven thee (Matt.9:2)
(Believe ye that I am able to do this? (Matt.9:28)
(The harvest truly is plenteous, but the labourers are few (Matt.9:37)
(Freely ye have received, freely give (Matt.10:8)
(Be ye therefore wise as serpents, & harmless as doves (Matt.10:16)
(What ye hear in the ear, that preach ye upon the housetops (Matt.10:27)
(Come unto Me, all ye that labour & are heavy laden, & I will give you rest (Matt.11:28)
(Learn of me; for I am meek & lowly in heart (Matt.11:29)
(My yoke is easy, & My burden is light (Matt.11:30)
(A bruised reed shall He not break (Matt.12:20)
(He that gathereth not with Me scattereth abroad (Matt.12:30)
(Then shall the righteous shine forth as the sun in the kingdom of their Father (Matt.13:43)
(The kingdom of heaven is like unto treasure hid in a field (Matt.13:44)
(I will build My church; & the gates of hell shall not prevail against it (Matt.16:18)
(Many that are first shall be last; & the last shall be first (Matt.19:30)
(He that shall endure unto the end, the same shall be saved (Matt.24:13)
(& this gospel of the kingdom shall be preached in all the world (Matt.24:14)
(Heaven & earth shall pass away, but My words shall not pass away (Matt.24:35)
(Who then is a faithful & wise servant (Matt.24:45)
(Watch & pray, that ye enter not into temptation (Matt.26:41)
(The spirit indeed is willing, but the flesh is weak (Matt.26:41)
(Go ye therefore, & teach all nations (Matt.28:19)
(&, lo, I am with you alway, even unto the end of the world (Matt.28:20)
(& in the morning, rising up a great while before day… prayed (Mark 1:35)
(& as many as touched Him were made whole (Mark 6:56)
(What would ye that I should do for you? (Mark 10:36)
(Repent ye, & believe the gospel (Mark 11:15)
(Go ye into all the world, & preach the gospel to every creature (Mark 16:15)
(Blessed are ye that weep now: for ye shall laugh (Luke 6:21)
(Take up his cross daily, & follow me (Luke 9:23)
(In the beginning was the Word, & the Word was with God, & the Word was God (John 1:1)
(& the light shineth in darkness; & the darkness comprehended it not. (John 1:5)
(But as many as received him, to them gave he power to become the sons of God (John 1:12)
(& the Word was made flesh, & dwelt among us (John 1:14)
(Make straight the way of the LORD (John 1:23)
(Hereafter ye shall see heaven open (John 1:51)
(Except a man be born again, he cannot see the kingdom of God (John 3:3)
(That which is born of the Spirit is spirit (John 3:6)
(The wind bloweth where it listeth, & thou hearest the sound thereof (John 3:8)
(For God so loved the world, that he gave his only begotten Son (John 3:16)
(He that doeth truth cometh to the light (John 3:21)
(He that believeth on the Son hath everlasting life (John 3:36)
(The true worshippers shall worship the Father in spirit & in truth (John 4:23)
(God is a Spirit: & they that worship Him must worship him in spirit & in truth (John 4:24)
(Jesus saith unto them, My meat is to do the will of Him that sent Me, & to finish His work. (John 4:34)
(He that reapeth receiveth wages, & gathereth fruit unto life eternal (John 4:36)
(He that soweth & he that reapeth may rejoice together (John 4:36)
(Behold, thou art made whole: sin no more (John 5:14)
(Search the Scriptures; for in them ye think ye have eternal life (John 5:39)
(Labour not for the meat which perisheth, but for that meat which endureth unto everlasting life (John 6:27)
(This is the work of God, that ye believe on Him whom He hath sent (John 6:29)
(For the bread of God is He which cometh down from heaven, & giveth life unto the world (John 6:33)
(He that believeth on Me shall never thirst (John 6:35)
(He that cometh to me shall never hunger (John 6:35)
(Him that cometh to Me I will in no wise cast out (John 6:37)
(He that believeth on Me hath everlasting life (John 6:47)
(It is the spirit that quickeneth; the flesh profiteth nothing (John 6:63)
(The words that I speak unto you, they are spirit, & they are life (John 6:63)
(If any man thirst, let him come unto me, & drink (John 7:37)
(He that believeth on Me… out of his belly shall flow rivers of living water (John 7:38)
(He that followeth Me shall not walk in darkness (John 8:12)
(If ye continue in My word, then are ye My disciples indeed (John 8:31)
(& ye shall know the truth, & the truth shall make you free (John 8:32)
(Whosoever committeth sin is the servant of sin (John 8:34)
(He that is of God heareth God's words (John 8:47)
(If a man keep My saying, he shall never see death (John 8:51)
(I am the door: by Me if any man enter in, he shall be saved. (John 10:9)
(I am come that they might have life, & that they might have it more abundantly (John 10:10)
(I am the good shepherd: the good shepherd giveth His life for the sheep (John 10:11)
(I am the good shepherd, & know My sheep, & am known of Mine (John 10:14)
(My sheep hear My voice, & I know them, & they follow Me. (John 10:27)
(Neither shall any man pluck them out of My hand (John 10:28)
(I am the resurrection, & the life (John 11:25)
(If thou wouldest believe, thou shouldest see the glory of God (John 11:40)
(If any man serve Me, him will My Father honour (John 12:26)
(If any man serve me, let him follow Me (John 12:26)
(While ye have light, believe in the light, that ye may be the children of light (John 12:36)
(A new comm&ment I give unto you, that ye love one another (John 13:34)
(By this shall all men know that ye are My disciples, if ye have love one to another (John 13:35)
(Let not your heart be troubled: ye believe in God, believe also in Me. (John 14:1)
(I will come again, & receive you unto Myself (John 14:3)
(I am the way, the truth, & the life (John 14:6)
(He that believeth on Me, the works that I do shall he do also (John 14:12)
(& whatsoever ye shall ask in My name, that will I do (John 14:13)
(If ye love Me, keep My commandments (John 14:15)
(The Comforter, which is the Holy Ghost… He shall teach you all things (John 14:26)
(Let not your heart be troubled, neither let it be afraid (John 14:27)
(Abide in Me, & I in you (John 15:4)
(He that abideth in Me, & I in him, the same bringeth forth much fruit (John 15:5)
(Herein is My Father glorified, that ye bear much fruit (John 15:8)
(I loved you: continue ye in My love (John 15:9)
(That My joy might remain in you, & that your joy might be full. (John 15:11)
(This is My comm&ment, that ye love one another, as I have loved you (John 15:12)
(Ye are My friends, if ye do whatsoever I comm& you. (John 15:14)
(Ye have not chosen Me, but I have chosen you (John 15:16)
(I have chosen you, & ordained you, that ye should go & bring forth fruit (John 15:16)
(Ask, & ye shall receive, that your joy may be full (John 16:24)
(Be of good cheer; I have overcome the world (John 16:33)
(As my Father hath sent Me, even so send I you (John 20:21)
(Blessed are they that have not seen, & yet have believed (John 20:29)
(Lord,… grant unto Thy servants, that with all boldness they may speak Thy word (Acts 4:29)
(The will of the Lord be done (Acts 21:14)
(I am not ashamed of the gospel of Christ (Rom.1:16)
(The just shall live by faith (Rom.1:17)
(The goodness of God leadeth thee to repentance (Rom.2:4)
(All have sinned, & come short of the glory of God (Rom.3:23)
(Blessed are they whose iniquities are forgiven (Rom.4:7)
(Being justified by faith, we have peace with God (Rom.5:1)
(The love of God is shed abroad in our hearts by the Holy Ghost (Rom.5:5)
(While we were yet sinners, Christ died for us (Rom.5:8)
(Being now justified by His blood, we shall be saved from wrath through Him (Rom.5:9)
(When we were enemies, we were reconciled to God by the death of His Son (Rom.5:10)
(Where sin abounded, grace did much more abound (Rom.5:20)
(Yield yourselves unto God, as those that are alive from the dead (Rom.6:13)
(The gift of God is eternal life through Jesus Christ our Lord (Rom.6:23)
(We should bring forth fruit unto God (Rom.7:4)
(We are debtors, not to the flesh, to live after the flesh (Rom.8:12)
(For as many as are led by the Spirit of God, they are the sons of God. (Rom.8:14)
(Likewise the Spirit also helpeth our infirmities. (Rom.8:26)
(All things work together for good to them that love God (Rom.8:28)
(If God be for us, who can be against us? (Rom.8:31)
(Who shall lay any thing to the charge of God's elect? It is God that justifieth. (Rom.8:33)
(Nay, in all these things we are more than conquerors through him that loved us. (Rom.8:37)
(Nor height, nor depth, nor any other creature, shall be able to separate us from the love of God (Rom.8:39)
(O man, who art thou that repliest against God? (Rom.9:20)
(Whosoever believeth on him shall not be ashamed (Rom.9:33)
(The word is nigh thee, even in thy mouth, & in thy heart (Rom.10:8)
(With the mouth confession is made unto salvation (Rom.10:10)
(The same Lord over all is rich unto all that call upon Him. (Rom.10:12)
(How beautiful are the feet of them that preach the gospel of peace (Rom.10:15)
(Faith cometh by hearing, & hearing by the word of God (Rom.10:17)
(The gifts & calling of God are without repentance (Rom.11:29)
(O the depth of the riches both of the wisdom & knowledge of God! (Rom.11:33)
(& be not conformed to this world (Rom.12:2)
(Transformed by the renewing of your mind (Rom.12:2)
(That ye may prove what is that good, & acceptable, & perfect, will of God (Rom.12:2)
(Let love be without dissimulation (Rom.12:9)
(Be kindly affectioned one to another with brotherly love (Rom.12:10)
(Fervent in spirit; serving the Lord (Rom.12:11)
(Rejoicing in hope; patient in tribulation; continuing instant in prayer (Rom.12:12)
(Rejoice with them that do rejoice, & weep with them that weep (Rom.12:15)
(Recompense to no man evil for evil (Rom.12:17)
(Live peaceably with all men (Rom.12:18)
(Be not overcome of evil, but overcome evil with good (Rom.12:21)
(Owe no man any thing, but to love one another (Rom.13:8)
(Now it is high time to awake out of sleep (Rom.13:11)
(Now is our salvation nearer than when we believed (Rom.13:11)
(Let us put on the armour of light (Rom.13:12)
(Put ye on the Lord Jesus Christ (Rom.13:14)
(So then every one of us shall give account of himself to God (Rom.14:12)
(The kingdom of God is… righteousness, & peace, & joy in the Holy Ghost. (Rom.14:17)
(Happy is he that condemneth not himself in that thing which he alloweth (Rom.14:22)
(Christ… received us to the glory of God (Rom.15:7)
(Now the God of hope fill you with all joy & peace in believing (Rom.15:13)
(The God of peace shall bruise Satan under your feet shortly (Rom.16:20)
(For the preaching of the cross is… the power of God (1Cor.1:18)
(We preach Christ crucified (1Cor.1:23)
(God hath chosen the weak things of the world to confound the things which are mighty (1Cor.1:27)
(But God hath chosen the foolish things of the world to confound the wise (1Cor.1:27)
(We have received, not the spirit of the world, but the spirit which is of God (1Cor.2:12)
(Every man shall receive his own reward according to his own labour (1Cor.3:8)
(We are labourers together with God (1Cor.3:9)
(Know ye not that ye are the temple of God? (1Cor.3:16)
(The temple of God is holy, which temple ye are (1Cor.3:17)
(Ye are Christ's; & Christ is God's (1Cor.3:23)
(We are fools for Christ's sake (1Cor.4:10)
(The kingdom of God is not in word, but in power (1Cor.4:20)
(For even Christ our passover is sacrificed for us (1Cor.5:7)
(He that is joined unto the Lord is one spirit (1Cor.6:17)
(Ye are bought with a price (1Cor.6:20)
(Glorify God in your body, & in your spirit, which are God's (1Cor.6:20)
(For the fashion of this world passeth away (1Cor.7:31)
(Knowledge puffeth up, but charity edifieth (1Cor.8:1)
(If any man love God, the same is known of Him (1Cor.8:3)
(So run, that ye may obtain (1Cor.9:24)
(Let him that thinketh he st&eth take heed lest he fall (1Cor.10:12)
(God is faithful, who will not suffer you to be tempted above that ye are able (1Cor.10:13)
(Do all to the glory of God (1Cor.10:31)
(The manifestation of the Spirit is given to every man (1Cor.12:7)
(Covet earnestly the best gifts (1Cor.12:31)
(Charity suffereth long, & is kind (1Cor.13:4)
(Charity vaunteth not itself, is not puffed up (1Cor.13:4)
(Charity never faileth (1Cor.13:8)
(& now abideth faith, hope, charity, these three (1Cor.13:13)
(Follow after charity, & desire spiritual gifts (1Cor.14:1)
(Christ died for our sins according to the Scriptures (1Cor.15:3)
(But now is Christ risen from the dead (1Cor.15:20)
(Evil communications corrupt good manners (1Cor.15:33)
(Thanks be to God, which giveth us the victory (1Cor.15:57)
(Be ye stedfast, unmoveable (1Cor.15:58)
(Always abounding in the work of the Lord (1Cor.15:58)
(Your labour is not in vain in the Lord (1Cor.15:58)
(Watch ye, st& fast in the faith, quit you like men, be strong (1Cor.16:13)
(Let all your things be done with charity (1Cor.16:14)
(All the promises of God in Him are yea, & in Him Amen (2Cor.1:20)
(He which stablisheth us with you in Christ, & hath anointed us, is God (2Cor.1:21)
(Thanks be unto God, which always causeth us to triumph in Christ (2Cor.2:14)
(We are unto God a sweet savour of Christ (2Cor.2:15)
(Ye are… the epistle of Christ (2Cor.3:3)
(Our sufficiency is of God (2Cor.3:5)
(Where the Spirit of the Lord is, there is liberty (2Cor.3:17)
(God, who comm&ed the light to shine out of darkness, hath shined in our hearts (2Cor.4:6)
(We have this treasure in earthen vessels (2Cor.4:7)
(We are troubled on every side, yet not distressed (2Cor.4:8)
(We are perplexed, but not in despair (2Cor.4:8)
(Our light affliction… worketh for us a far more exceeding & eternal weight of glory (2Cor.4:17)
(The things which are seen are temporal; but the things which are not seen are eternal (2Cor.4:18)
(We have a building of God… in the heavens (2Cor.5:1)
(We walk by faith, not by sight (2Cor.5:7)
(If any man be in Christ, he is a new creature (2Cor.5:17)
(Old things are passed away; behold, all things are become new (2Cor.5:17)
(Be ye reconciled to God (2Cor.5:20)
(He hath made him to be sin for us, who knew no sin (2Cor.5:21)
(Behold, now is the accepted time; behold, now is the day of salvation (2Cor.6:2)
(As sorrowful, yet alway rejoicing (2Cor.6:10)
(Touch not the unclean thing; & I will receive you (2Cor.6:17)
(& will be a Father unto you, & ye shall be my sons & daughters (2Cor.6:18)
(Having therefore these promises, dearly beloved, let us cleanse ourselves (2Cor.7:1)
(For godly sorrow worketh repentance to salvation (2Cor.7:10)
(Yet for your sakes He became poor, that ye through His poverty might be rich (2Cor.8:9)
(Providing for honest things, not only in the sight of the Lord, but also in the sight of men (2Cor.8:21)
(He which soweth sparingly shall reap also sparingly (2Cor.9:6)
(God loveth a cheerful giver (2Cor.9:7)
(God is able to make all grace abound toward you (2Cor.9:8)
(That ye… may abound to every good work (2Cor.9:8)
(He that ministereth seed to the sower… multiply your seed sown (2Cor.9:10)
(Thanks be unto God for his unspeakable gift (2Cor.9:15)
(For though we walk in the flesh, we do not war after the flesh. (2Cor.10:3)
(The weapons of our warfare are not carnal (2Cor.10:4)
(Bringing into captivity every thought to the obedience of Christ (2Cor.10:5)
(But he that glorieth, let him glory in the LORD (2Cor.10:17)
(My strength is made perfect in weakness (2Cor.12:9)
(Be perfect, be of good comfort, be of one mind, live in peace. (2Cor.13:11)
(Be of one mind, live in peace (2Cor.13:11)
(Who gave Himself for our sins (Gal.1:4)
(God accepteth no man's person (Gal.2:6)
(We have believed in Jesus Christ, that we might be justified by the faith of Christ (Gal.2:16)
(Nevertheless I live; yet not I, but Christ liveth in me (Gal.2:20)
(They which are of faith, the same are the children of Abraham (Gal.3:7)
(Christ hath redeemed us from the curse of the law (Gal.3:13)
(If ye be Christ's, then are ye Abraham's seed, & heirs according to the promise (Gal.3:29)
(God hath sent forth the Spirit of his Son into your hearts (Gal.4:6)
(Wherefore thou art no more a servant, but a son; & if a son, then an heir of God (Gal.4:7)
(St& fast therefore in the liberty wherewith Christ hath made us free (Gal.5:1)
(By love serve one another (Gal.5:13)
(Thou shalt love thy neighbour as thyself (Gal.5:14)
(Walk in the Spirit (Gal.5:16)
(The fruit of the Spirit is love, joy, peace… (Gal.5:22)
(They that are Christ's have crucified the flesh (Gal.5:24)
(If we live in the Spirit, let us also walk in the Spirit (Gal.5:25)
(Bear ye one another's burdens, & so fulfil the law of Christ. (Gal.6:2)
(Let every man prove his own work (Gal.6:4)
(Whatsoever a man soweth, that shall he also reap (Gal.6:7)
(He that soweth to the Spirit shall of the Spirit reap life everlasting (Gal.6:8)
(Let us not be weary in well doing (Gal.6:9)
(As we have therefore opportunity, let us do good unto all men (Gal.6:10)
(Blessed be the God,… who hath blessed us with every spiritual blessing (Eph.1:3)
(He hath chosen us in Him before the foundation of the world (Eph.1:4)
(In whom we have redemption through His blood, the forgiveness of sins (Eph.1:7)
(Ye were sealed with that holy Spirit of promise (Eph.1:13)
(What the riches of the glory of His inheritance in the saints (Eph.1:18)
(What is the exceeding greatness of His power to us (Eph.1:19)
(For by grace are ye saved through faith (Eph.2:8)
(For we are his workmanship, created in Christ Jesus unto good works (Eph.2:10)
(Ye who sometimes were far off are made nigh by the blood of Christ (Eph.2:13)
(Ye are… fellowcitizens with the saints, & of the household of God (Eph.2:19)
(I beseech you that ye walk worthy of the vocation wherewith ye are called (Eph.4:1)
(Unto every one of us is given grace according to the measure of the gift of Christ. (Eph.4:7)
(Speak every man truth with his neighbour (Eph.4:25)
(Grieve not the holy Spirit of God (Eph.4:30)
(Be ye kind one to another, tenderhearted (Eph.4:32)
(Forgiving one another, even as God for Christ's sake hath forgiven you (Eph.4:32)
(Be ye therefore followers of God, as dear children (Eph.5:1)
(Christ also hath loved us, & hath given Himself for us (Eph.5:2)
(Walk as children of light (Eph.5:8)
(The fruit of the Spirit is in all goodness & righteousness & truth (Eph.5:9)
(Proving what is acceptable unto the Lord (Eph.5:10)
(See then that ye walk circumspectly, not as fools, but as wise (Eph.5:15)
(Redeeming the time, because the days are evil (Eph.5:16)
(Underst&ing what the will of the Lord is (Eph.5:17)
(Be filled with the Spirit (Eph.5:18)
(Christ also loved the church, & gave Himself for it (Eph.5:25)
(Be strong in the Lord, & in the power of his might (Eph.6:10)
(Put on the whole armour of God (Eph.6:11)
(We wrestle not against flesh & blood (Eph.6:12)
(St& therefore, having your loins girt about with truth (Eph.6:14)
(Praying always with all prayer & supplication in the Spirit (Eph.6:18)
(For to me to live is Christ, & to die is gain (Phil.1:21)
(Only let your conversation be as it becometh the gospel of Christ (Phil.1:27)
(Let nothing be done through strife or vainglory (Phil.2:3)
(In lowliness of mind let each esteem other better than themselves (Phil.2:3)
(Work out your own salvation with fear & trembling (Phil.2:12)
(It is God which worketh in you both to will & to do of his good pleasure (Phil.2:13)
(Do all things without murmurings & disputings (Phil.2:14)
(I count all things but loss for the excellency of the knowledge of Christ Jesus (Phil.3:8)
(Forgetting those things which are behind, & reaching forth unto those things which are before (Phil.3:13)
(Our conversation is in heaven (Phil.3:20)
(Rejoice in the Lord alway: & again I say, Rejoice! (Phil.4:4)
(Let your moderation be known unto all men (Phil.4:5)
(The Lord is at hand (Phil.4:5)
(Be careful for nothing (Phil.4:6)
(Let your requests be made known unto God (Phil.4:6)
(I am instructed both to be full & to be hungry, both to abound & to suffer need (Phil.4:12)
(I can do all things through Christ which strengtheneth me (Phil.4:13)
(God shall supply all your need (Phil.4:19)
(That ye might be filled with the knowledge of His will (Col.1:9)
(That ye might walk worthy of the Lord (Col.1:10)
(He hath delivered us from the power of darkness (Col.1:13)
(In whom we have redemption through His blood, even the forgiveness of sins (Col.1:14)
(Christ in you, the hope of glory (Col.1:27)
(As ye have therefore received Christ Jesus the Lord, so walk ye in Him (Col.2:6)
(In Him dwelleth all the fulness of the Godhead bodily (Col.2:9)
(Ye are complete in Him (Col.2:10)
(Hath he quickened together with Him, having forgiven you all trespasses (Col.2:13)
(Seek those things which are above, where Christ sitteth on the right hand of God (Col.3:1)
(Set your affection on things above, not on things on the earth (Col.3:2)
(Your life is hid with Christ in God (Col.3:3)
(& above all these things put on charity (Col.3:14)
(Charity… is the bond of perfectness (Col.3:14)
(Let the peace of God rule in your hearts (Col.3:15)
(Be ye thankful (Col.3:15)
(Let the word of Christ dwell in you richly (Col.3:16)
(Do all in the name of the Lord Jesus (Col.3:17)
(& whatsoever ye do, do it heartily, as to the Lord (Col.3:23)
(Of the Lord ye shall receive the reward of the inheritance (Col.3:24)
(Continue in prayer (Col.4:2)
(Let your speech be alway with grace, seasoned with salt (Col.4:6)
(Take heed to the ministry which thou hast received in the Lord, that thou fulfil it (Col.4:17)
(& the Lord make you to increase & abound in love one toward another (1Thess.3:12)
(For this is the will of God, even your sanctification (1Thess.4:3)
(That every one of you should know how to possess his vessel in sanctification & honour (1Thess.4:4)
(For God hath not called us unto uncleanness, but unto holiness (1Thess.4:7)
(That ye may walk honestly toward them that are without (1Thess.4:12)
(Ye are all the children of light, & the children of the day (1Thess.5:5)
(Let us not sleep, as do others; but let us watch & be sober (1Thess.5:6)
(Let us, who are of the day, be sober (1Thess.5:8)
(For God hath not appointed us to wrath, but to obtain salvation (1Thess.5:9)
(Ever follow that which is good (1Thess.5:15)
(Rejoice evermore (1Thess.5:16)
(Pray without ceasing (1Thess.5:17)
(In every thing give thanks (1Thess.5:18)
(Prove all things; hold fast that which is good (1Thess.5:21)
(The very God of peace sanctify you wholly (1Thess.5:23)
(I pray God your whole spirit & soul & body be preserved blameless (1Thess.5:23)
(God hath from the beginning chosen you to salvation (2Thess.2:13)
(The Lord is faithful, who shall stablish you, & keep you from evil (2Thess.3:3)
(If any would not work, neither should he eat (2Thess.3:10)
(The Lord of peace himself give you peace always by all means (2Thess.3:16)
(I thank Christ Jesus our Lord, who hath enabled me (1Tim.1:12)
(Christ Jesus came into the world to save sinners (1Tim.1:15)
(There is one God, & one mediator between God & men (1Tim.2:5)
(Great is the mystery of godliness: God was manifest in the flesh (1Tim.3:16)
(We trust in the living God, who is the Saviour of all men (1Tim.4:10)
(Till I come, give attendance to reading, to exhortation, to doctrine. (1Tim.4:13)
(Neglect not the gift that is in thee, which was given thee (1Tim.4:14)
(Take heed unto thyself, & unto the doctrine; continue in them. (1Tim.4:16)
(Thou shalt not muzzle the ox that treadeth out the corn (1Tim.5:18)
(The labourer is worthy of his reward (1Tim.5:18)
(Keep thyself pure (1Tim.5:22)
(Godliness with contentment is great gain (1Tim.6:6)
(Having food & raiment let us be therewith content (1Tim.6:8)
(The love of money is the root of all evil (1Tim.6:10)
(Follow after righteousness, godliness, faith, love, patience, meekness (1Tim.6:11)
(Fight the good fight of faith (1Tim.6:12)
(Lay hold on eternal life, whereunto thou art also called (1Tim.6:12)
(Keep that which is committed to thy trust (1Tim.6:20)
(I put thee in remembrance that thou stir up the gift of God, which is in thee (2Tim.1:6)
(God hath not given us the spirit of fear; but of power, & of love, & of a sound mind (2Tim.1:7)
(Be not thou therefore ashamed of the testimony of our Lord (2Tim.1:8)
(Hold fast the form of sound words (2Tim.1:13)
(That good thing which was committed unto thee keep by the Holy Ghost, who dwelleth in us. (2Tim.1:14)
(Be strong in the grace that is in Christ Jesus (2Tim.2:1)
(Thou therefore endure hardness, as a good soldier of Jesus Christ (2Tim.2:3)
(The husb&man that laboureth must be first partaker of the fruits. (2Tim.2:6)
(Remember that Jesus Christ of the seed of David was raised from the dead (2Tim.2:8)
(If we suffer, we shall also reign with Him (2Tim.2:12)
(The Lord knoweth them that are His (2Tim.2:19)
(Follow righteousness, faith, charity, peace, with them that call on the Lord (2Tim.2:22)
(The servant of the Lord must not strive; but be gentle unto all men (2Tim.2:24)
(All that will live godly in Christ Jesus shall suffer persecution (2Tim.3:12)
(Continue thou in the things which thou hast learned (2Tim.3:14)
(All Scripture is given by inspiration of God, & is profitable for doctrine (2Tim.3:16)
(That the man of God may be perfect, throughly furnished unto all good works (2Tim.3:17)
(Preach the word; be instant in season, out of season. (2Tim.4:2)
(But watch thou in all things (2Tim.4:5)
(Do the work of an evangelist, make full proof of thy ministry. (2Tim.4:5)
(In all things shewing thyself a pattern of good works (Titus 2:7)
(The grace of God that bringeth salvation hath appeared to all men (Titus 2:11)
(Not by works of righteousness which we have done, but according to His mercy He saved us (Titus 3:5)
(That the communication of thy faith may become effectual (Philem.1:6)
(Having confidence in thy obedience I wrote unto thee (Philem.1:21)
(We ought to give the more earnest heed to the things which we have heard (Heb.2:1)
(For both He that sanctifieth & they who are sanctified are all of one (Heb.2:11)
(As the children are partakers of flesh & blood, he also himself likewise took part of the same (Heb.2:14)
(He is able to succour them that are tempted (Heb.2:18)
(For every house is builded by some man; but he that built all things is God. (Heb.3:4)
(Christ as a son over His own house; whose house are we (Heb.3:6)
(Exhort one another daily (Heb.3:13)
(For we are made partakers of Christ (Heb.3:14)
(We which have believed do enter into rest (Heb.4:3)
(Let us labour therefore to enter into that rest (Heb.4:11)
(The word of God is quick, & powerful, & sharper than any twoedged sword (Heb.4:12)
(All things are naked & opened unto the eyes of Him with whom we have to do. (Heb.4:13)
(Let us hold fast our profession (Heb.4:14)
(Let us therefore come boldly unto the throne of grace (Heb.4:16)
(Yet learned He obedience by the things which He suffered (Heb.5:8)
(Leaving the principles of the doctrine of Christ, let us go on unto perfection (Heb.6:1)
(For God is not unrighteous to forget your work & labour of love (Heb.6:10)
(Surely blessing I will bless thee, & multiplying I will multiply thee (Heb.6:14)
(Jesus… is able also to save them to the uttermost that come unto God by Him (Heb.7:25)
(Jesus… ever liveth to make intercession (Heb.7:25)
(I will be to them a God, & they shall be to me a people (Heb.8:10)
(I will put my laws into their mind, & write them in their hearts (Heb.8:10)
(The blood of Christ… purge your conscience from dead works (Heb.9:14)
(We are sanctified through the offering of the body of Jesus Christ (Heb.10:10)
(Their sins & iniquities will I remember no more (Heb.10:17)
(Let us hold fast the profession of our faith without wavering (Heb.10:23)
(Let us consider one another to provoke unto love & to good works (Heb.10:24)
(Not forsaking the assembling of ourselves together (Heb.10:25)
(Cast not away therefore your confidence, which hath great recompence of reward (Heb.10:35)
(Ye have need of patience, that, after ye have done the will of God, ye might receive the promise (Heb.10:36)
(Yet a little while, & he that shall come will come, & will not tarry (Heb.10:37)
(We are… of them that believe to the saving of the soul (Heb.10:39)
(Faith is the substance of things hoped for, the evidence of things not seen. (Heb.11:1)
(But without faith it is impossible to please Him (Heb.11:6)
(He that cometh to God must believe that He is (Heb.11:6)
(God is a rewarder of them that diligently seek him (Heb.11:6)
(Strangers & pilgrims on the earth (Heb.11:13)
(Through faith subdued kingdoms, wrought righteousness, obtained promises (Heb.11:33)
(God having provided some better thing for us (Heb.11:40)
(Let us lay aside every weight, & the sin which doth so easily beset us (Heb.12:1)
(Let us run with patience the race that is set before us (Heb.12:1)
(For the joy that was set before Him endured the cross (Heb.12:2)
(Consider Him that endured such contradiction of sinners against Himself (Heb.12:3)
(Ye have not yet resisted unto blood, striving against sin (Heb.12:4)
(Nor faint when thou art rebuked of Him (Heb.12:5)
(Lift up the hands which hang down, & the feeble knees (Heb.12:12)
(Follow peace with all men, & holiness (Heb.12:14)
(Ye are come unto mount Sion, & unto the city of the living God (Heb.12:22)
(Wherefore we receiving a kingdom which cannot be moved, let us have grace (Heb.12:28)
(Let us have grace, whereby we may serve God acceptably with reverence & godly fear (Heb.12:28)
(Let brotherly love continue. (Heb.13:1)
(Be not forgetful to entertain strangers (Heb.13:2)
(Be content with such things as ye have (Heb.13:5)
(I will never leave thee, nor forsake thee (Heb.13:5)
(The LORD is my helper, & I will not fear what man shall do unto me (Heb.13:6)
(Jesus Christ the same yesterday, & to day, & for ever (Heb.13:8)
(It is a good thing that the heart be established with grace (Heb.13:9)
(Jesus, that he might sanctify the people with His own blood, suffered without the gate (Heb.13:12)
(Let us go forth therefore unto Him without the camp, bearing His reproach. (Heb.13:13)
(Here have we no continuing city, but we seek one to come (Heb.13:14)
(By Him let us offer the sacrifice of praise to God continually, that is, the fruit of our lips (Heb.13:15)
(To do good & to communicate forget not (Heb.13:16)
(Obey them that have the rule over you (Heb.13:17)
(He that wavereth is like a wave of the sea (James 1:6)
(Blessed is the man that endureth temptation (James 1:12)
(Every good gift & every perfect gift is from above (James 1:17)
(Of His own will begat he us with the word of truth (James 1:18)
(Receive with meekness the engrafted word (James 1:21)
(Be ye doers of the word (James 1:22)
(To keep himself unspotted from the world (James 1:27)
(Mercy rejoiceth against judgment. (James 2:13)
(By works was faith made perfect (James 2:22)
(Faith without works is dead (James 2:26)
(If any man offend not in word, the same is a perfect man (James 3:2)
(The spirit that dwelleth in us lusteth to envy (James 4:5)
(God resisteth the proud, but giveth grace unto the humble (James 4:6)
(Submit yourselves therefore to God. Resist the devil. (James 4:7)
(Draw nigh to God, & He will draw nigh to you. (James 4:8)
(Humble yourselves in the sight of the Lord, & He shall lift you up. (James 4:10)
(Be patient therefore, brethren, unto the coming of the Lord (James 5:7)
(Be ye also patient; stablish your hearts (James 5:8)
(The coming of the Lord draweth nigh (James 5:8)
(That the Lord is very pitiful, & of tender mercy (James 5:11)
(Is any among you afflicted? let him pray. (James 5:13)
(Is any merry? let him sing psalms. (James 5:13)
(The prayer of faith shall save the sick (James 5:15)
(Pray one for another, that ye may be healed (James 5:16)
(The effectual fervent prayer of a righteous man availeth much (James 5:16)
(Blessed be the God… which hath begotten us again unto a lively hope (1Pet.1:3)
(Be ye holy in all manner of conversation (1Pet.1:15)
(Be ye holy; for I am holy (1Pet.1:16)
(Ye were not redeemed with corruptible things, as silver & gold, from your vain conversation (1Pet.1:18)
(Love one another with a pure heart fervently (1Pet.1:22)
(The word of the Lord endureth for ever (1Pet.1:25)
(As newborn babes, desire the sincere milk of the word (1Pet.2:2)
(Ye have tasted that the Lord is gracious (1Pet.2:3)
(He that believeth on him shall not be confounded (1Pet.2:6)
(But ye are a chosen generation, a royal priesthood, an holy nation (1Pet.2:9)
(Honour all men. Love the brotherhood. Fear God. Honour the king. (1Pet.2:17)
(Who His own self bare our sins in His own body on the tree (1Pet.2:24)
(By whose stripes ye were healed (1Pet.2:24)
(Ye were as sheep going astray; but are now returned unto the Shepherd (1Pet.2:25)
(The eyes of the Lord are over the righteous (1Pet.3:12)
(If ye suffer for righteousness' sake, happy are ye (1Pet.3:14)
(Sanctify the Lord God in your hearts (1Pet.3:15)
(Christ hath once suffered for sins,… that He might bring us to God (1Pet.3:18)
(The end of all things is at hand (1Pet.4:7)
(Use hospitality one to another without grudging (1Pet.4:9)
(As every man hath received the gift, even so minister the same one to another (1Pet.4:10)
(If any man speak, let him speak as the oracles of God (1Pet.4:11)
(If any man minister, let him do it as of the ability which God giveth (1Pet.4:11)
(The spirit of glory & of God resteth upon you (1Pet.4:14)
(Ye shall receive a crown of glory that fadeth not away (1Pet.5:4)
(Humble yourselves therefore under the mighty hand of God, that He may exalt you in due time (1Pet.5:6)
(Casting all your care upon him; for he careth for you (1Pet.5:7)
(As His divine power hath given unto us all things that pertain unto life & godliness (2Pet.1:3)
(Exceeding great & precious promises (2Pet.1:4)
(Give diligence to make your calling & election sure (2Pet.1:10)
(The Lord knoweth how to deliver the godly out of temptations (2Pet.2:9)
(The Lord is not slack concerning His promise (2Pet.3:9)
(The day of the Lord will come as a thief in the night (2Pet.3:10)
(Be diligent that ye may be found of him in peace, without spot, & blameless (2Pet.3:14)
(Account that the longsuffering of our Lord is salvation (2Pet.3:15)
(Grow in grace, & in the knowledge of our Lord (2Pet.3:18)
(For the life was manifested (1John 1:2)
(Our fellowship is with the Father, & with His Son Jesus Christ (1John 1:3)
(That your joy may be full (1John 1:4)
(God is light, & in him is no darkness at all (1John 1:5)
(The blood of Jesus Christ his Son cleanseth us from all sin (1John 1:7)
(We have an advocate with the Father, Jesus Christ the righteous (1John 2:1)
(Whoso keepeth His word, in him verily is the love of God perfected (1John 2:5)
(The darkness is past, & the true light now shineth (1John 2:8)
(Your sins are forgiven you for His name's sake (1John 2:12)
(& the world passeth away, & the lust thereof (1John 2:17)
(He that doeth the will of God abideth for ever (1John 2:17)
(This is the promise that he hath promised us, even eternal life (1John 2:25)
(& now, little children, abide in Him (1John 2:28)
(Behold, what manner of love the Father hath bestowed upon us (1John 3:1)
(Beloved, now are we the sons of God (1John 3:2)
(He was manifested to take away our sins (1John 3:5)
(Hereby perceive we the love of God, because He laid down His life for us (1John 3:16)
(Let us not love in word, neither in tongue; but in deed & in truth (1John 3:18)
(If our heart condemn us not, then have we confidence toward God (1John 3:21)
(Greater is he that is in you, than he that is in the world (1John 4:4)
(Let us love one another: for love is of God (1John 4:7)
(God is love (1John 4:8)
(The love of God was manifested toward us (1John 4:9)
(He loved us, & sent his Son to be the propitiation for our sins (1John 4:10)
(if God so loved us, we ought also to love one another (1John 4:11)
(The Father sent the Son to be the Saviour of the world (1John 4:14)
(& we have known & believed the love that God hath to us (1John 4:16)
(There is no fear in love; but perfect love casteth out fear (1John 4:18)
(We love Him, because he first loved us (1John 4:19)
(His commandments are not grievous (1John 5:3)
(Whatsoever is born of God overcometh the world (1John 5:4)
(This is the victory that overcometh the world, even our faith. (1John 5:4)
(God hath given to us eternal life (1John 5:11)
(That ye may know that ye have eternal life (1John 5:13)
(He that is begotten of God keepeth himself, & that wicked one toucheth him not (1John 5:18)
(We know that we are of God, & the whole world lieth in wickedness (1John 5:19)
(The Son of God is come, & hath given us an underst&ing (1John 5:20)
(This is love, that we walk after His commandments (2John 1:6)
(I wish above all things that thou mayest prosper & be in health (3John 1:2)
(Keep yourselves in the love of God (Jude 1:21)
(Blessed is he that readeth, & they that hear the words of this prophecy (Rev.1:3)
(Behold, He cometh with clouds; & every eye shall see Him (Rev.1:7)
(I am Alpha & Omega, the beginning & the ending, saith the Lord (Rev.1:8)
(I know thy works, & thy labour, & thy patience (Rev.2:2)
(To him that overcometh will I give to eat of the tree of life (Rev.2:7)
(Be thou faithful unto death, & I will give thee a crown of life (Rev.2:10)
(To him that overcometh will I give to eat of the hidden manna (Rev.2:17)
(Remember therefore how thou hast received & heard, & hold fast, & repent (Rev.3:3)
(Behold, I have set before thee an open door, & no man can shut it (Rev.3:8)
(Behold, I come quickly! (Rev.3:11)
(Hold that fast which thou hast, that no man take thy crown (Rev.3:11)
(Him that overcometh will I make a pillar in the temple of My God (Rev.3:12)
(He that hath an ear, let him hear what the Spirit saith unto the churches (Rev.3:13)
(As many as I love, I rebuke & chasten (Rev.3:19)
(Behold, I st& at the door, & knock (Rev.3:20)
(If any man hear my voice, & open the door, I will come in to him (Rev.3:20)
(To him that overcometh will I grant to sit with Me in My throne (Rev.3:21)
(Holy, holy, holy, LORD God Almighty (Rev.4:8)
(Thou art worthy, O Lord, to receive glory & honour & power (Rev.4:11)
(& hast made us unto our God kings & priests (Rev.5:10)
(Worthy is the Lamb that was slain (Rev.5:12)
(Salvation to our God which sitteth upon the throne, & unto the Lamb! (Rev.7:10)
(& God shall wipe away all tears from their eyes (Rev.7:17)
(& the smoke of the incense, which came with the prayers of the saints, ascended up before God (Rev.8:4)
(They overcame him by the blood of the Lamb, & by the word of their testimony (Rev.12:11)
(Worship Him that made heaven, & earth, & the sea, & the fountains of waters (Rev.14:7)
(Great & marvellous are Thy works, Lord God Almighty! (Rev.15:3)
(Who shall not fear Thee, O Lord, & glorify Thy name? (Rev.15:4)
(Thou art righteous, O Lord, which art, & wast, & shalt be (Rev.16:5)
(Even so, Lord God Almighty, true & righteous are thy judgments (Rev.16:7)
(Blessed is he that watcheth, & keepeth his garments (Rev.16:15)
(Behold, I make all things new (Rev.21:5)
(I will give unto him that is athirst of the fountain of the water of life freely (Rev.21:6)
(He that overcometh shall inherit all things (Rev.21:7)
(Blessed is he that keepeth the sayings of the prophecy of this book (Rev.22:7)
(He that is holy, let him be holy still (Rev.22:11)
(Whosoever will, let him take the water of life freely (Rev.22:17)
(Even so, come, Lord Jesus! (Rev.22:20)

    ` + bcolors.ENDC
    fmt.Println()
    for _, s := range scriptures {
        fmt.Print(string(s))
        time.Sleep(45 * time.Millisecond)
    }
}

func CommandMents() {
     fmt.Printf(bcolors.GREEN + `
                   ABOUT BIBLE VERSE DESKTOP   ` + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + `
[ #.(vrsoft.org ).........................................]
[ #name.......King James Version..........................]
[ #copyright....Public Domain.............................]
[ #language....english....................................]
[ ........................................................]
[ Scripture taken from the King James Version.............]
[ Public Domain...........................................]
` + bcolors.ENDC)
    fmt.Printf(bcolors.GREEN + `
[ What is there 4 U 2 gain the whole world & loose your...]
[ soul? Be smart your Creator has good plans for you......]
[.....Life Tip.: Defeat the devil by fasting & praying....]
` + bcolors.ENDC)
    fmt.Printf(bcolors.YELLOW + `
                         ¯\_(ツ)_/¯
    ` + bcolors.ENDC)
    fmt.Print(bcolors.GREEN + `
                 📜 The Ten commandments.` + bcolors.ENDC)
    fmt.Print(bcolors.BLUE + `
[ 1. You shall have no other gods before Me...............]
[ 2. You shall make no idols..............................]
[ 3. You shall not take the name of the Lord your God in..]
[  . vein.................................................]
[ 4. Keep the Sabbath day holy............................]
[ 5. Honor your father and your mother....................]
[ 6. You shall not murder.................................]
[ 7. You shall not commit adultery........................]
[ 8. You shall not steal..................................]
[ 9. You shall not bear false witness against your eighbor]
[ 10 You shall not covet..................................]
[.........................................................]` + bcolors.ENDC)
    fmt.Printf(bcolors.YELLOW + `
             Type 0. To.Exit & Go To Main Menu

` + bcolors.ENDC)
}

func (b *Scriptures) TheWord() {
    rand.Seed(time.Now().UnixNano())
    menu := rand.Intn(1006) + 1

    switch menu {
    case 1:
        bcolors.Colors(`(📜In the beginning God created the heaven & the  (Gen.1:1)` + bcolors.ENDC)
    case 2:
        bcolors.Colors(`(📜& God saw the light, that it was good: & God d (Gen.1:4)` + bcolors.ENDC)
    case 3:
        bcolors.Colors(`(📜I will bless thee… & thou shalt be a blessing (Gen.12:2)` + bcolors.ENDC)
    case 4:
        bcolors.Colors(`(📜I am thy shield, & thy exceeding great reward (Gen.15:1)` + bcolors.ENDC)
    case 5:
        bcolors.Colors(`(📜Fear ye not, st& still, & see the salvation ..(Ex.14:13)` + bcolors.ENDC)
    case 6:
        bcolors.Colors(`(📜I will make all My goodness pass before thee..(Ex.33:19)` + bcolors.ENDC)
    case 7:
        bcolors.Colors(`(📜The LORD God, merciful & gracious .............(Ex.34:6)` + bcolors.ENDC)
    case 8:
        bcolors.Colors(`(📜I set my tabernacle among you ...............(Lev.26:11)` + bcolors.ENDC)
    case 9:
        bcolors.Colors(`(📜I will walk among you, & will be your God ...(Lev.26:12)` + bcolors.ENDC)
    case 10:
        bcolors.Colors(`(📜The LORD is longsuffering, & of great mercy .(Num.14:18)` + bcolors.ENDC)
    case 11:
        bcolors.Colors(`(📜Thou shalt love the LORD thy God with all thin(Deut.6:5)` + bcolors.ENDC)
    case 12:
        bcolors.Colors(`(📜Shall ye lay up these my words in your heart(Deut.11:18)` + bcolors.ENDC)
    case 13:
        bcolors.Colors(`(📜Thou shalt rejoice before the LORD thy God .(Deut.12:18)` + bcolors.ENDC)
    case 14:
        bcolors.Colors(`(📜Blessed shalt thou be in the city, & blesse..(Deut.28:3)` + bcolors.ENDC)
    case 15:
        bcolors.Colors(`(📜Blessed shall be the fruit of thy body, & the(Deut.28:4)` + bcolors.ENDC)
    case 16:
        bcolors.Colors(`(📜Blessed shall be thy basket & thy store .....(Deut.28:5)` + bcolors.ENDC)
    case 17:
        bcolors.Colors(`(📜Blessed shalt thou be when thou comest in, & (Deut.28:6)` + bcolors.ENDC)
    case 18:
        bcolors.Colors(`(📜They shall come out against thee one way, & .(Deut.28:7)` + bcolors.ENDC)
    case 19:
        bcolors.Colors(`(📜& the LORD shall make thee the head, & not t(Deut.28:13)` + bcolors.ENDC)
    case 20:
        bcolors.Colors(`(📜Be strong & of a good courage, fear not, nor (Deut.31:6)` + bcolors.ENDC)
    case 21:
        bcolors.Colors(`(📜I will not fail thee, nor forsake thee .......(Josh.1:5)` + bcolors.ENDC)
    case 22:
        bcolors.Colors(`(📜Only be thou strong & very courageous ........(Josh.1:7)` + bcolors.ENDC)
    case 23:
        bcolors.Colors(`(📜This Book of the Law shall not depart out of t(Josh.1:8)` + bcolors.ENDC)
    case 24:
        bcolors.Colors(`(📜Have not I comm&ed thee? Be strong & of a ....(Josh.1:9)` + bcolors.ENDC)
    case 25:
        bcolors.Colors(`(📜Be not afraid, neither be thou dismayed: for t(Josh.1:9)` + bcolors.ENDC)
    case 26:
        bcolors.Colors(`(📜Sanctify yourselves: for to morrow the LORD w (Josh.3:5)` + bcolors.ENDC)
    case 27:
        bcolors.Colors(`(📜The LORD your God, He it is that fighteth fo(Josh.23:10)` + bcolors.ENDC)
    case 28:
        bcolors.Colors(`(📜Nay; but we will serve the LORD! ...........(Josh.24:21)` + bcolors.ENDC)
    case 29:
        bcolors.Colors(`(📜I will never break my covenant with you ......(Judg.2:1)` + bcolors.ENDC)
    case 30:
        bcolors.Colors(`(📜Them that love him be as the sun when he goe (Judg.5:31)` + bcolors.ENDC)
    case 31:
        bcolors.Colors(`(📜Thy people shall be my people, & thy God my G(Ruth 1:16)` + bcolors.ENDC)
    case 32:
        bcolors.Colors(`(📜He will keep the feet of His saints ..........(1Sam.2:9)` + bcolors.ENDC)
    case 33:
        bcolors.Colors(`(📜Only fear the LORD, & serve Him in truth wi (1Sam.12:24)` + bcolors.ENDC)
    case 34:
        bcolors.Colors(`(📜I will shew thee what thou shalt do .........(1Sam.16:3)` + bcolors.ENDC)
    case 35:
        bcolors.Colors(`(📜Man looketh on the outward appearance, but t (1Sam.16:7)` + bcolors.ENDC)
    case 36:
        bcolors.Colors(`(📜Yet he hath made with me an everlasting cove (2Sam.23:5)` + bcolors.ENDC)
    case 37:
        bcolors.Colors(`(📜Give therefore thy servant an underst&ing h ..(1Kin.3:9)` + bcolors.ENDC)
    case 38:
        bcolors.Colors(`(📜The LORD your God ye shall fear; & he sha ..(2Kin.17:39)` + bcolors.ENDC)
    case 39:
        bcolors.Colors(`(📜Serve Him with a perfect heart & with a will (1Chr.28:9)` + bcolors.ENDC)
    case 40:
        bcolors.Colors(`(📜The LORD searcheth all hearts, & underst&e ..(1Chr.28:9)` + bcolors.ENDC)
    case 41:
        bcolors.Colors(`(📜If thou seek him, he will be found of thee ..(1Chr.28:9)` + bcolors.ENDC)
    case 42:
        bcolors.Colors(`(📜The LORD is able to give thee much more than (2Chr.25:9)` + bcolors.ENDC)
    case 43:
        bcolors.Colors(`(📜Will not turn away His face from you, if ye  (2Chr.30:9)` + bcolors.ENDC)
    case 44:
        bcolors.Colors(`(📜He did it with all his heart, & prospered ..(2Chr.31:21)` + bcolors.ENDC)
    case 45:
        bcolors.Colors(`(📜The hand of our God is upon all them for go ..(Ezdra 8:22)` + bcolors.ENDC)
    case 46:
        bcolors.Colors(`(📜The joy of the LORD is your strength..........(Neh.8:10)` + bcolors.ENDC)
    case 47:
        bcolors.Colors(`(📜& who knoweth whether thou art come to t...(Esther.4:14)` + bcolors.ENDC)
    case 48:
        bcolors.Colors(`(📜Till he fill thy mouth with laughing, & thy l (Job 8:21)` + bcolors.ENDC)
    case 49:
        bcolors.Colors(`(📜I know that my redeemer liveth ..............(Job 19:25)` + bcolors.ENDC)
    case 50:
        bcolors.Colors(`(📜Blessed is the man that walketh not in the couns(Ps.1:1)` + bcolors.ENDC)
    case 51:
        bcolors.Colors(`(📜The LORD knoweth the way of the righteous ......(Ps.1:6)` + bcolors.ENDC)
    case 52:
        bcolors.Colors(`(📜Serve the LORD with fear, & rejoice with tremb (Ps.2:11)` + bcolors.ENDC)
    case 53:
        bcolors.Colors(`(📜Blessed are all they that put their trust in H (Ps.2:12)` + bcolors.ENDC)
    case 54:
        bcolors.Colors(`(📜But thou, O LORD, art a shield for me, my glory (Ps.3:3)` + bcolors.ENDC)
    case 55:
        bcolors.Colors(`(📜Lead me, O LORD, in thy righteousness ..........(Ps.5:8)` + bcolors.ENDC)
    case 56:
        bcolors.Colors(`(📜For thou, LORD, wilt bless the righteous ......(Ps.5:12)` + bcolors.ENDC)
    case 57:
        bcolors.Colors(`(📜For the righteous God trieth the hearts ........(Ps.7:9)` + bcolors.ENDC)
    case 58:
        bcolors.Colors(`(📜I will praise thee, O LORD, with my whole heart (Ps.9:1)` + bcolors.ENDC)
    case 59:
        bcolors.Colors(`(📜For thou, LORD, hast not forsaken them that se (Ps.9:10)` + bcolors.ENDC)
    case 60:
        bcolors.Colors(`(📜The LORD is King for ever & ever .............(Ps.10:16)` + bcolors.ENDC)
    case 61:
        bcolors.Colors(`(📜LORD, thou hast heard the desire of the humbl (Ps.10:17)` + bcolors.ENDC)
    case 62:
        bcolors.Colors(`(📜The LORD is the portion of mine inheritance &..(Ps.16:5)` + bcolors.ENDC)
    case 63:
        bcolors.Colors(`(📜Hide me under the shadow of thy wings .........(Ps.17:8)` + bcolors.ENDC)
    case 64:
        bcolors.Colors(`(📜I will love thee, O LORD, my strength .........(Ps.18:1)` + bcolors.ENDC)
    case 65:
        bcolors.Colors(`(📜The LORD is my rock, & my fortress, & my deliv (Ps.18:2)` + bcolors.ENDC)
    case 66:
        bcolors.Colors(`(📜For thou wilt light my c&le ..................(Ps.18:28)` + bcolors.ENDC)
    case 67:
        bcolors.Colors(`(📜The LORD my God will enlighten my darkness ...(Ps.18:28)` + bcolors.ENDC)
    case 68:
        bcolors.Colors(`(📜It is God that girdeth me with strength ......(Ps.18:32)` + bcolors.ENDC)
    case 69:
        bcolors.Colors(`(📜The LORD liveth; & blessed be my rock ........(Ps.18:46)` + bcolors.ENDC)
    case 70:
        bcolors.Colors(`(📜Let the God of my salvation be exalted .......(Ps.18:46)` + bcolors.ENDC)
    case 71:
        bcolors.Colors(`(📜The heavens declare the glory of God; & the fi (Ps.19:1)` + bcolors.ENDC)
    case 72:
        bcolors.Colors(`(📜The law of the LORD is perfect, converting the (Ps.19:7)` + bcolors.ENDC)
    case 73:
        bcolors.Colors(`(📜The testimony of the LORD is sure, making wise (Ps.19:7)` + bcolors.ENDC)
    case 74:
        bcolors.Colors(`(📜The statutes of the LORD are right, rejoicing  (Ps.19:8)` + bcolors.ENDC)
    case 75:
        bcolors.Colors(`(📜The comm&ment of the LORD is pure, enlighten ..(Ps.19:8)` + bcolors.ENDC)
    case 76:
        bcolors.Colors(`(📜The fear of the LORD is clean, enduring for ev (Ps.19:9)` + bcolors.ENDC)
    case 77:
        bcolors.Colors(`(📜The judgments of the LORD are true & righteous (Ps.19:9)` + bcolors.ENDC)
    case 78:
        bcolors.Colors(`(📜Some trust in chariots, & some in horses: but  (Ps.20:7)` + bcolors.ENDC)
    case 79:
        bcolors.Colors(`(📜They shall praise the LORD that seek Him .....(Ps.22:26)` + bcolors.ENDC)
    case 80:
        bcolors.Colors(`(📜All the ends of the world shall remember & tu (Ps.22:27)` + bcolors.ENDC)
    case 81:
        bcolors.Colors(`(📜A seed shall serve Him .......................(Ps.22:30)` + bcolors.ENDC)
    case 82:
        bcolors.Colors(`(📜The LORD is my shepherd; I shall not want .....(Ps.23:1)` + bcolors.ENDC)
    case 83:
        bcolors.Colors(`(📜The LORD strong & mighty, the LORD mighty in b (Ps.24:8)` + bcolors.ENDC)
    case 84:
        bcolors.Colors(`(📜O my God, I trust in Thee: let me not be asham (Ps.25:2)` + bcolors.ENDC)
    case 85:
        bcolors.Colors(`(📜Shew me thy ways, O LORD; teach me Thy paths. .(Ps.25:4)` + bcolors.ENDC)
    case 86:
        bcolors.Colors(`(📜Good & upright is the LORD: therefore will He  (Ps.25:8)` + bcolors.ENDC)
    case 87:
        bcolors.Colors(`(📜The secret of the LORD is with them that fear (Ps.25:14)` + bcolors.ENDC)
    case 88:
        bcolors.Colors(`(📜Mine eyes are ever toward the LORD; for He sh (Ps.25:15)` + bcolors.ENDC)
    case 89:
        bcolors.Colors(`(📜Teach me Thy way, O LORD .....................(Ps.27:11)` + bcolors.ENDC)
    case 90:
        bcolors.Colors(`(📜Wait on the LORD: be of good courage, & He sh (Ps.27:14)` + bcolors.ENDC)
    case 91:
        bcolors.Colors(`(📜The LORD is my strength & my shield ...........(Ps.28:7)` + bcolors.ENDC)
    case 92:
        bcolors.Colors(`(📜My heart trusted in Him, & I am helped ........(Ps.28:7)` + bcolors.ENDC)
    case 93:
        bcolors.Colors(`(📜The voice of the LORD is powerful; the voice o (Ps.29:4)` + bcolors.ENDC)
    case 94:
        bcolors.Colors(`(📜The LORD will give strength unto His people ..(Ps.29:11)` + bcolors.ENDC)
    case 95:
        bcolors.Colors(`(📜The LORD will bless his people with peace ....(Ps.29:11)` + bcolors.ENDC)
    case 96:
        bcolors.Colors(`(📜In his favour is life .........................(Ps.30:5)` + bcolors.ENDC)
    case 97:
        bcolors.Colors(`(📜Be of good courage, & he shall strengthen your(Ps.31:24)` + bcolors.ENDC)
    case 98:
        bcolors.Colors(`(📜Blessed is he whose transgression is forgiven .(Ps.32:1)` + bcolors.ENDC)
    case 99:
        bcolors.Colors(`(📜I will instruct thee & teach thee in the way w (Ps.32:8)` + bcolors.ENDC)
    case 100:
        bcolors.Colors(`(📜He that trusteth in the LORD, mercy shall com (Ps.32:10)` + bcolors.ENDC)
    case 101:
        bcolors.Colors(`(📜Be glad in the LORD, & rejoice ...............(Ps.32:11)` + bcolors.ENDC)
    case 102:
        bcolors.Colors(`(📜For the word of the LORD is right; & all His w (Ps.33:4)` + bcolors.ENDC)
    case 103:
        bcolors.Colors(`(📜The earth is full of the goodness of the LORD .(Ps.33:5)` + bcolors.ENDC)
    case 104:
        bcolors.Colors(`(📜Our soul waiteth for the LORD: He is our help (Ps.33:20)` + bcolors.ENDC)
    case 105:
        bcolors.Colors(`(📜They looked unto Him, & were lightened ........(Ps.34:5)` + bcolors.ENDC)
    case 106:
        bcolors.Colors(`(📜O taste & see that the LORD is good ...........(Ps.34:8)` + bcolors.ENDC)
    case 107:
        bcolors.Colors(`(📜They that seek the LORD shall not want any go (Ps.34:10)` + bcolors.ENDC)
    case 108:
        bcolors.Colors(`(📜The LORD is nigh unto them that are of a brok (Ps.34:18)` + bcolors.ENDC)
    case 109:
        bcolors.Colors(`(📜& my soul shall be joyful in the LORD: it shal (Ps.35:9)` + bcolors.ENDC)
    case 110:
        bcolors.Colors(`(📜Trust in the LORD, & do good ..................(Ps.37:3)` + bcolors.ENDC)
    case 111:
        bcolors.Colors(`(📜Dwell in the l&, & verily thou shalt be fed ...(Ps.37:3)` + bcolors.ENDC)
    case 112:
        bcolors.Colors(`(📜Delight thyself also in the LORD; & He shall g (Ps.37:4)` + bcolors.ENDC)
    case 113:
        bcolors.Colors(`(📜Commit thy way unto the LORD; trust also in Hi (Ps.37:5)` + bcolors.ENDC)
    case 114:
        bcolors.Colors(`(📜Rest in the LORD, & wait patiently for Him ....(Ps.37:7)` + bcolors.ENDC)
    case 115:
        bcolors.Colors(`(📜The meek shall inherit the earth; & shall del (Ps.37:11)` + bcolors.ENDC)
    case 116:
        bcolors.Colors(`(📜Wait on the LORD, & keep His way .............(Ps.37:34)` + bcolors.ENDC)
    case 117:
        bcolors.Colors(`(📜The salvation of the righteous is of the LORD (Ps.37:39)` + bcolors.ENDC)
    case 118:
        bcolors.Colors(`(📜Lord, all my desire is before thee ............(Ps.38:9)` + bcolors.ENDC)
    case 119:
        bcolors.Colors(`(📜Blessed is that man that maketh the LORD his t (Ps.40:4)` + bcolors.ENDC)
    case 120:
        bcolors.Colors(`(📜I delight to do Thy will, O my God ............(Ps.40:8)` + bcolors.ENDC)
    case 121:
        bcolors.Colors(`(📜Be still, & know that I am God ...............(Ps.46:10)` + bcolors.ENDC)
    case 122:
        bcolors.Colors(`(📜O clap your hands, all ye people; shout unto God (Ps.47:1)` + bcolors.ENDC)
    case 123:
        bcolors.Colors(`(📜The redemption of their soul is precious ......(Ps.49:9)` + bcolors.ENDC)
    case 124:
        bcolors.Colors(`(📜Our God shall come, & shall not keep silence ..(Ps.50:3)` + bcolors.ENDC)
    case 125:
        bcolors.Colors(`(📜Offer unto God thanksgiving; & pay thy vows u (Ps.50:14)` + bcolors.ENDC)
    case 126:
        bcolors.Colors(`(📜Call upon me in the day of trouble: I will de (Ps.50:15)` + bcolors.ENDC)
    case 127:
        bcolors.Colors(`(📜Whoso offereth praise glorifieth Me ..........(Ps.50:23)` + bcolors.ENDC)
    case 128:
        bcolors.Colors(`(📜The fool hath said in his heart, "There is no  (Ps.53:1)` + bcolors.ENDC)
    case 129:
        bcolors.Colors(`(📜Behold, God is mine helper: the Lord is with t (Ps.54:4)` + bcolors.ENDC)
    case 130:
        bcolors.Colors(`(📜Cast thy burden upon the LORD, & He shall sus (Ps.55:22)` + bcolors.ENDC)
    case 131:
        bcolors.Colors(`(📜He shall never suffer the righteous to be mov (Ps.55:22)` + bcolors.ENDC)
    case 132:
        bcolors.Colors(`(📜Verily there is a reward for the righteous ...(Ps.58:11)` + bcolors.ENDC)
    case 133:
        bcolors.Colors(`(📜God is my defence, & the God of my mercy .....(Ps.59:17)` + bcolors.ENDC)
    case 134:
        bcolors.Colors(`(📜Truly my soul waiteth upon God: from Him comet (Ps.62:1)` + bcolors.ENDC)
    case 135:
        bcolors.Colors(`(📜God is a refuge for us ........................(Ps.62:8)` + bcolors.ENDC)
    case 136:
        bcolors.Colors(`(📜God hath spoken once; twice have I heard this (Ps.62:11)` + bcolors.ENDC)
    case 137:
        bcolors.Colors(`(📜The righteous shall be glad in the LORD, & sh (Ps.64:10)` + bcolors.ENDC)
    case 138:
        bcolors.Colors(`(📜Thy God hath comm&ed thy strength ............(Ps.68:28)` + bcolors.ENDC)
    case 139:
        bcolors.Colors(`(📜But it is good for me to draw near to God ....(Ps.73:28)` + bcolors.ENDC)
    case 140:
        bcolors.Colors(`(📜Open thy mouth wide, & I will fill it ........(Ps.81:10)` + bcolors.ENDC)
    case 141:
        bcolors.Colors(`(📜The LORD God is a sun & shield ...............(Ps.84:11)` + bcolors.ENDC)
    case 142:
        bcolors.Colors(`(📜The LORD will give grace & glory .............(Ps.84:11)` + bcolors.ENDC)
    case 143:
        bcolors.Colors(`(📜The LORD shall give that which is good; & our (Ps.85:12)` + bcolors.ENDC)
    case 144:
        bcolors.Colors(`(📜For thou, Lord, art good, & ready to forgive;  (Ps.86:5)` + bcolors.ENDC)
    case 145:
        bcolors.Colors(`(📜He that dwelleth in the secret place of the Mo (Ps.91:1)` + bcolors.ENDC)
    case 146:
        bcolors.Colors(`(📜Thou shalt not be afraid for the terror by nig (Ps.91:5)` + bcolors.ENDC)
    case 147:
        bcolors.Colors(`(📜Lest thou dash thy foot against a stone ......(Ps.91:12)` + bcolors.ENDC)
    case 148:
        bcolors.Colors(`(📜The righteous shall flourish like the palm tr (Ps.92:12)` + bcolors.ENDC)
    case 149:
        bcolors.Colors(`(📜The LORD reigneth, He is clothed with majesty .(Ps.93:1)` + bcolors.ENDC)
    case 150:
        bcolors.Colors(`(📜The LORD knoweth the thoughts of man, that th (Ps.94:11)` + bcolors.ENDC)
    case 151:
        bcolors.Colors(`(📜The LORD will not cast off His people ........(Ps.94:14)` + bcolors.ENDC)
    case 152:
        bcolors.Colors(`(📜Give unto the LORD the glory due unto His name (Ps.96:8)` + bcolors.ENDC)
    case 153:
        bcolors.Colors(`(📜Worship the LORD in the Salvation of holiness .(Ps.96:9)` + bcolors.ENDC)
    case 154:
        bcolors.Colors(`(📜The LORD reigneth: the world also shall be es (Ps.96:10)` + bcolors.ENDC)
    case 155:
        bcolors.Colors(`(📜Light is sown for the righteous, & gladness f (Ps.97:11)` + bcolors.ENDC)
    case 156:
        bcolors.Colors(`(📜The LORD hath made known His salvation ........(Ps.98:2)` + bcolors.ENDC)
    case 157:
        bcolors.Colors(`(📜Serve the LORD with gladness .................(Ps.100:2)` + bcolors.ENDC)
    case 158:
        bcolors.Colors(`(📜I will set no wicked thing before mine eyes ..(Ps.101:3)` + bcolors.ENDC)
    case 159:
        bcolors.Colors(`(📜They shall perish, but thou shalt endure ....(Ps.102:26)` + bcolors.ENDC)
    case 160:
        bcolors.Colors(`(📜Bless the LORD, O my soul, & forget not all H (Ps.103:2)` + bcolors.ENDC)
    case 161:
        bcolors.Colors(`(📜Forget not all His benefits ..................(Ps.103:2)` + bcolors.ENDC)
    case 162:
        bcolors.Colors(`(📜Who forgiveth all thine iniquities; who heale (Ps.103:3)` + bcolors.ENDC)
    case 163:
        bcolors.Colors(`(📜As the heaven is high above the earth, so gr (Ps.103:11)` + bcolors.ENDC)
    case 164:
        bcolors.Colors(`(📜As far as the east is from the west, so far  (Ps.103:12)` + bcolors.ENDC)
    case 165:
        bcolors.Colors(`(📜I will sing unto the LORD as long as I live .(Ps.104:33)` + bcolors.ENDC)
    case 166:
        bcolors.Colors(`(📜Let the heart of them rejoice that seek the L (Ps.105:3)` + bcolors.ENDC)
    case 167:
        bcolors.Colors(`(📜Seek the LORD, & His strength: seek his face  (Ps.105:4)` + bcolors.ENDC)
    case 168:
        bcolors.Colors(`(📜O give thanks unto the LORD; for He is good:  (Ps.106:1)` + bcolors.ENDC)
    case 169:
        bcolors.Colors(`(📜He sent his word, & healed them .............(Ps.107:20)` + bcolors.ENDC)
    case 170:
        bcolors.Colors(`(📜Through God we shall do valiantly ...........(Ps.108:13)` + bcolors.ENDC)
    case 171:
        bcolors.Colors(`(📜I will greatly praise the LORD with my mouth (Ps.109:30)` + bcolors.ENDC)
    case 172:
        bcolors.Colors(`(📜He raiseth up the poor out of the dust, & lif (Ps.113:7)` + bcolors.ENDC)
    case 173:
        bcolors.Colors(`(📜Ye that fear the LORD, trust in the LORD: He (Ps.115:11)` + bcolors.ENDC)
    case 174:
        bcolors.Colors(`(📜He will bless them that fear the LORD .......(Ps.115:13)` + bcolors.ENDC)
    case 175:
        bcolors.Colors(`(📜I will take the cup of salvation, & call upo (Ps.116:13)` + bcolors.ENDC)
    case 176:
        bcolors.Colors(`(📜His merciful kindness is great toward us .....(Ps.117:2)` + bcolors.ENDC)
    case 177:
        bcolors.Colors(`(📜It is better to trust in the LORD than to put (Ps.118:8)` + bcolors.ENDC)
    case 178:
        bcolors.Colors(`(📜The LORD is my strength & song, & is become  (Ps.118:14)` + bcolors.ENDC)
    case 179:
        bcolors.Colors(`(📜The right hand of the LORD doeth valiantly ....(Ps.118:16)` + bcolors.ENDC)
    case 180:
        bcolors.Colors(`(📜This is the day which the LORD hath made; we (Ps.118:24)` + bcolors.ENDC)
    case 181:
        bcolors.Colors(`(📜Thy word is a lamp unto my feet, & a light  (Ps.119:105)` + bcolors.ENDC)
    case 182:
        bcolors.Colors(`(📜Great peace have they which love thy law ...(Ps.119:165)` + bcolors.ENDC)
    case 183:
        bcolors.Colors(`(📜My help cometh from the LORD, which made heav (Ps.121:2)` + bcolors.ENDC)
    case 184:
        bcolors.Colors(`(📜He will not suffer thy foot to be moved ......(Ps.121:3)` + bcolors.ENDC)
    case 185:
        bcolors.Colors(`(📜The LORD is thy keeper .......................(Ps.121:5)` + bcolors.ENDC)
    case 186:
        bcolors.Colors(`(📜I was glad when they said unto me, "Let us go (Ps.122:1)` + bcolors.ENDC)
    case 187:
        bcolors.Colors(`(📜Our soul is escaped as a bird out of the snar (Ps.124:7)` + bcolors.ENDC)
    case 188:
        bcolors.Colors(`(📜The snare is broken, & we are escaped ........(Ps.124:7)` + bcolors.ENDC)
    case 189:
        bcolors.Colors(`(📜Our help is in the name of the LORD ..........(Ps.124:8)` + bcolors.ENDC)
    case 190:
        bcolors.Colors(`(📜They that trust in the LORD shall be as Mount (Ps.125:1)` + bcolors.ENDC)
    case 191:
        bcolors.Colors(`(📜They that sow in tears shall reap in joy .....(Ps.126:5)` + bcolors.ENDC)
    case 192:
        bcolors.Colors(`(📜Lift up your hands… & bless the LORD ...........(Ps.134:2)` + bcolors.ENDC)
    case 193:
        bcolors.Colors(`(📜The LORD will perfect that which concerneth m (Ps.138:8)` + bcolors.ENDC)
    case 194:
        bcolors.Colors(`(📜The LORD upholdeth all that fall ............(Ps.145:14)` + bcolors.ENDC)
    case 195:
        bcolors.Colors(`(📜The LORD is righteous in all His ways .......(Ps.145:17)` + bcolors.ENDC)
    case 196:
        bcolors.Colors(`(📜The LORD is nigh unto all them that call upo (Ps.145:18)` + bcolors.ENDC)
    case 197:
        bcolors.Colors(`(📜He healeth the broken in heart, & bindeth up  (Ps.147:3)` + bcolors.ENDC)
    case 198:
        bcolors.Colors(`(📜The LORD lifteth up the meek .................(Ps.147:6)` + bcolors.ENDC)
    case 199:
        bcolors.Colors(`(📜Let the children of Zion be joyful in their K (Ps.149:2)` + bcolors.ENDC)
    case 200:
        bcolors.Colors(`(📜For the LORD taketh pleasure in His people ...(Ps.149:4)` + bcolors.ENDC)
    case 201:
        bcolors.Colors(`(📜Let every thing that hath breath praise the L (Ps.150:6)` + bcolors.ENDC)
    case 202:
        bcolors.Colors(`(📜The fear of the LORD is the beginning of know (Prov.1:7)` + bcolors.ENDC)
    case 203:
        bcolors.Colors(`(📜My son, if sinners entice thee, consent thou (Prov.1:10)` + bcolors.ENDC)
    case 204:
        bcolors.Colors(`(📜I will pour out my spirit unto you, I will m (Prov.1:23)` + bcolors.ENDC)
    case 205:
        bcolors.Colors(`(📜For the LORD giveth wisdom: out of His mouth  (Prov.2:6)` + bcolors.ENDC)
    case 206:
        bcolors.Colors(`(📜He layeth up sound wisdom for the righteous ..(Prov.2:7)` + bcolors.ENDC)
    case 207:
        bcolors.Colors(`(📜He keepeth the paths of judgment .............(Prov.2:8)` + bcolors.ENDC)
    case 208:
        bcolors.Colors(`(📜Walk in the way of good men, & keep the path (Prov.2:20)` + bcolors.ENDC)
    case 209:
        bcolors.Colors(`(📜He blesseth the habitation of the just ......(Prov.3:33)` + bcolors.ENDC)
    case 210:
        bcolors.Colors(`(📜I give you good doctrine, forsake ye not my l (Prov.4:2)` + bcolors.ENDC)
    case 211:
        bcolors.Colors(`(📜Get wisdom: & with all thy getting get under. (Prov.4:7)` + bcolors.ENDC)
    case 212:
        bcolors.Colors(`(📜I have taught thee in the way of wisdom; I h (Prov.4:11)` + bcolors.ENDC)
    case 213:
        bcolors.Colors(`(📜When thou goest, thy steps shall not be stra (Prov.4:12)` + bcolors.ENDC)
    case 214:
        bcolors.Colors(`(📜Take fast hold of instruction ...............(Prov.4:13)` + bcolors.ENDC)
    case 215:
        bcolors.Colors(`(📜The path of the just is as the shining light (Prov.4:18)` + bcolors.ENDC)
    case 216:
        bcolors.Colors(`(📜Attend to my words; incline thine ear unto m (Prov.4:20)` + bcolors.ENDC)
    case 217:
        bcolors.Colors(`(📜My words… are life unto those that find them (Prov.4:22)` + bcolors.ENDC)
    case 218:
        bcolors.Colors(`(📜Keep thy heart with all diligence ...........(Prov.4:23)` + bcolors.ENDC)
    case 219:
        bcolors.Colors(`(📜Ponder the path of thy feet, & let all thy w (Prov.4:26)` + bcolors.ENDC)
    case 220:
        bcolors.Colors(`(📜Drink waters out of thine own cistern, & run (Prov.5:15)` + bcolors.ENDC)
    case 221:
        bcolors.Colors(`(📜For the ways of man are before the eyes of t (Prov.5:21)` + bcolors.ENDC)
    case 222:
        bcolors.Colors(`(📜For the comm&ment is a lamp; & the law is li (Prov.6:23)` + bcolors.ENDC)
    case 223:
        bcolors.Colors(`(📜My son, keep my words ........................(Prov.7:1)` + bcolors.ENDC)
    case 224:
        bcolors.Colors(`(📜Keep my commandments, & live; & my law as the a (Prov.7:2)` + bcolors.ENDC)
    case 225:
        bcolors.Colors(`(📜Wisdom is better than rubies ................(Prov.8:11)` + bcolors.ENDC)
    case 226:
        bcolors.Colors(`(📜Those that seek me early shall find me ......(Prov.8:17)` + bcolors.ENDC)
    case 227:
        bcolors.Colors(`(📜Blessed are they that keep my ways ..........(Prov.8:32)` + bcolors.ENDC)
    case 228:
        bcolors.Colors(`(📜The LORD will not suffer the soul of the rig (Prov.10:3)` + bcolors.ENDC)
    case 229:
        bcolors.Colors(`(📜The hand of the diligent maketh rich ..........(Prov.10:4)` + bcolors.ENDC)
    case 230:
        bcolors.Colors(`(📜The blessing of the LORD, it maketh rich, & (Prov.10:22)` + bcolors.ENDC)
    case 231:
        bcolors.Colors(`(📜The desire of the righteous shall be grante (Prov.10:24)` + bcolors.ENDC)
    case 232:
        bcolors.Colors(`(📜The hope of the righteous shall be gladness (Prov.10:28)` + bcolors.ENDC)
    case 233:
        bcolors.Colors(`(📜By the blessing of the upright the city is  (Prov.11:11)` + bcolors.ENDC)
    case 234:
        bcolors.Colors(`(📜He that watereth shall be watered also hims (Prov.11:25)` + bcolors.ENDC)
    case 235:
        bcolors.Colors(`(📜He that tilleth his l& shall be satisfied w (Prov.12:11)` + bcolors.ENDC)
    case 236:
        bcolors.Colors(`(📜He that keepeth his mouth keepeth his life ..(Prov.13:3)` + bcolors.ENDC)
    case 237:
        bcolors.Colors(`(📜The light of the righteous rejoiceth ........(Prov.13:9)` + bcolors.ENDC)
    case 238:
        bcolors.Colors(`(📜He that walketh with wise men shall be wise (Prov.13:20)` + bcolors.ENDC)
    case 239:
        bcolors.Colors(`(📜The wisdom of the prudent is to understand h (Prov.14:8)` + bcolors.ENDC)
    case 240:
        bcolors.Colors(`(📜The tabernacle of the upright shall flouri. (Prov.14:11)` + bcolors.ENDC)
    case 241:
        bcolors.Colors(`(📜A true witness delivereth souls ............(Prov.14:25)` + bcolors.ENDC)
    case 242:
        bcolors.Colors(`(📜The fear of the LORD is a fountain of life .(Prov.14:27)` + bcolors.ENDC)
    case 243:
        bcolors.Colors(`(📜A sound heart is the life of the flesh .....(Prov.14:30)` + bcolors.ENDC)
    case 244:
        bcolors.Colors(`(📜A soft answer turneth away wrath ............(Prov.15:1)` + bcolors.ENDC)
    case 245:
        bcolors.Colors(`(📜The eyes of the LORD are in every place .....(Prov.15:3)` + bcolors.ENDC)
    case 246:
        bcolors.Colors(`(📜He that refuseth instruction despiseth his  (Prov.15:32)` + bcolors.ENDC)
    case 247:
        bcolors.Colors(`(📜The fear of the LORD is the instruction of  (Prov.15:33)` + bcolors.ENDC)
    case 248:
        bcolors.Colors(`(📜Before honour is humility ..................(Prov.15:33)` + bcolors.ENDC)
    case 249:
        bcolors.Colors(`(📜Commit thy works unto the LORD, & thy though (Prov.16:3)` + bcolors.ENDC)
    case 250:
        bcolors.Colors(`(📜He that ruleth his spirit than he that take.(Prov.16:32)` + bcolors.ENDC)
    case 251:
        bcolors.Colors(`(📜The furnace for gold: but the LORD trieth th.(Prov.17:3)` + bcolors.ENDC)
    case 252:
        bcolors.Colors(`(📜A merry heart doeth good like a medicine....(Prov.17:22)` + bcolors.ENDC)
    case 253:
        bcolors.Colors(`(📜A broken spirit drieth the bones............(Prov.17:22)` + bcolors.ENDC)
    case 254:
        bcolors.Colors(`(📜He that hath knowledge spareth his words....(Prov.17:27)` + bcolors.ENDC)
    case 256:
        bcolors.Colors(`(📜The name of the LORD is a strong tower......(Prov.18:10)` + bcolors.ENDC)
    case 257:
        bcolors.Colors(`(📜Death & life are in the power of the tongue.(Prov.18:21)` + bcolors.ENDC)
    case 258:
        bcolors.Colors(`(📜A man that hath friends must shew himself f.(Prov.18:24)` + bcolors.ENDC)
    case 259:
        bcolors.Colors(`(📜The discretion of a man deferreth his anger.(Prov.19:11)` + bcolors.ENDC)
    case 260:
        bcolors.Colors(`(📜He that keepeth the comm&ment keepeth his o.(Prov.19:16)` + bcolors.ENDC)
    case 261:
        bcolors.Colors(`(📜He that hath pity upon the poor lendeth unt.(Prov.19:17)` + bcolors.ENDC)
    case 262:
        bcolors.Colors(`(📜It is an honour for a man to cease from stri.(Prov.20:3)` + bcolors.ENDC)
    case 263:
        bcolors.Colors(`(📜Wait on the LORD, & he shall save thee......(Prov.20:22)` + bcolors.ENDC)
    case 264:
        bcolors.Colors(`(📜The LORD pondereth the hearts................(Prov.21:2)` + bcolors.ENDC)
    case 265:
        bcolors.Colors(`(📜A good name is rather to be chosen than grea.(Prov.22:1)` + bcolors.ENDC)
    case 266:
        bcolors.Colors(`(📜He that hath a bountiful eye shall be blesse.(Prov.22:9)` + bcolors.ENDC)
    case 267:
        bcolors.Colors(`(📜Drowsiness shall clothe a man with rags.....(Prov.23:21)` + bcolors.ENDC)
    case 268:
        bcolors.Colors(`(📜Through wisdom is an house builded; & by und.(Prov.24:3)` + bcolors.ENDC)
    case 269:
        bcolors.Colors(`(📜By wise counsel thou shalt make thy war......(Prov.24:6)` + bcolors.ENDC)
    case 270:
        bcolors.Colors(`(📜In multitude of counsellors there is safety..(Prov.24:6)` + bcolors.ENDC)
    case 271:
        bcolors.Colors(`(📜For a just man falleth seven times, & riset.(Prov.24:16)` + bcolors.ENDC)
    case 272:
        bcolors.Colors(`(📜A soft tongue breaketh the bone.............(Prov.25:15)` + bcolors.ENDC)
    case 273:
        bcolors.Colors(`(📜The righteous are bold as a lion.............(Prov.28:1)` + bcolors.ENDC)
    case 274:
        bcolors.Colors(`(📜When righteous men do rejoice, there is gre.(Prov.28:12)` + bcolors.ENDC)
    case 275:
        bcolors.Colors(`(📜He that tilleth his l& shall have plenty of.(Prov.28:19)` + bcolors.ENDC)
    case 276:
        bcolors.Colors(`(📜The righteous considereth the cause of the p.(Prov.29:7)` + bcolors.ENDC)
    case 277:
        bcolors.Colors(`(📜Honour shall uphold the humble in spirit....(Prov.29:23)` + bcolors.ENDC)
    case 278:
        bcolors.Colors(`(📜To every thing there is a season, & a time to (Eccl.3:1)` + bcolors.ENDC)
    case 279:
        bcolors.Colors(`(📜A time to cast away stones, & a time to gathe.(Eccl.3:5)` + bcolors.ENDC)
    case 280:
        bcolors.Colors(`(📜A time to keep silence, & a time to speak.....(Eccl.3:7)` + bcolors.ENDC)
    case 281:
        bcolors.Colors(`(📜Whatsoever God doeth, it shall be for ever...(Eccl.3:14)` + bcolors.ENDC)
    case 282:
        bcolors.Colors(`(📜A threefold cord is not quickly broken.......(Eccl.4:12)` + bcolors.ENDC)
    case 283:
        bcolors.Colors(`(📜Keep thy foot when thou goest to the house of.(Eccl.5:1)` + bcolors.ENDC)
    case 284:
        bcolors.Colors(`(📜Seeing there be many things that increase va.(Eccl.6:11)` + bcolors.ENDC)
    case 285:
        bcolors.Colors(`(📜Anger resteth in the bosom of fools...........(Eccl.7:9)` + bcolors.ENDC)
    case 286:
        bcolors.Colors(`(📜God hath made man upright; but they have sou.(Eccl.7:29)` + bcolors.ENDC)
    case 287:
        bcolors.Colors(`(📜A man's wisdom maketh his face to shine.......(Eccl.8:1)` + bcolors.ENDC)
    case 288:
        bcolors.Colors(`(📜A wise man's heart discerneth both time & jud.(Eccl.8:5)` + bcolors.ENDC)
    case 289:
        bcolors.Colors(`(📜It shall be well with them that fear God.....(Eccl.8:12)` + bcolors.ENDC)
    case 290:
        bcolors.Colors(`(📜Let thy head lack no ointment.................(Eccl.9:8)` + bcolors.ENDC)
    case 291:
        bcolors.Colors(`(📜Whatsoever thy hand findeth to do, do it with..(Eccl.9:10)` + bcolors.ENDC)
    case 292:
        bcolors.Colors(`(📜Wisdom is better than weapons of war.........(Eccl.9:18)` + bcolors.ENDC)
    case 293:
        bcolors.Colors(`(📜He that diggeth a pit shall fall into........(Eccl.10:8)` + bcolors.ENDC)
    case 294:
        bcolors.Colors(`(📜Whoso breaketh an hedge, a serpent shall bit.(Eccl.10:8)` + bcolors.ENDC)
    case 295:
        bcolors.Colors(`(📜Cast thy bread upon the waters...............(Eccl.11:1)` + bcolors.ENDC)
    case 296:
        bcolors.Colors(`(📜He that observeth the wind shall not sow.....(Eccl.11:4)` + bcolors.ENDC)
    case 297:
        bcolors.Colors(`(📜Remember now thy Creator in the days of thy..(Eccl.12:1)` + bcolors.ENDC)
    case 298:
        bcolors.Colors(`(📜Vanity of vanities, saith the Preacher; all..(Eccl.12:8)` + bcolors.ENDC)
    case 299:
        bcolors.Colors(`(📜Fear God, & keep his commandments: for this.(Eccl.12:13)` + bcolors.ENDC)
    case 300:
        bcolors.Colors(`(📜Take us the foxes, the little foxes, that sp.(Song 2:15)` + bcolors.ENDC)
    case 301:
        bcolors.Colors(`(📜Many waters cannot quench love................(Song 8:7)` + bcolors.ENDC)
    case 302:
        bcolors.Colors(`(📜Learn to do well...............................(Is.1:17)` + bcolors.ENDC)
    case 303:
        bcolors.Colors(`(📜Though your sins be as scarlet, they shall be..(Is.1:18)` + bcolors.ENDC)
    case 304:
        bcolors.Colors(`(📜He will teach us of his ways, & we will walk in.(Is.2:3)` + bcolors.ENDC)
    case 305:
        bcolors.Colors(`(📜Whom shall I send, & who will go for Us?........(Is.6:8)` + bcolors.ENDC)
    case 306:
        bcolors.Colors(`(📜Then said I, Here am I; send me.................(Is.6:8)` + bcolors.ENDC)
    case 307:
        bcolors.Colors(`(📜The LORD JEHOVAH is my strength & my song......(Is.12:2)` + bcolors.ENDC)
    case 308:
        bcolors.Colors(`(📜In the LORD JEHOVAH is everlasting strength....(Is.26:4)` + bcolors.ENDC)
    case 309:
        bcolors.Colors(`(📜My people shall dwell in a peaceable habitati.(Is.32:18)` + bcolors.ENDC)
    case 310:
        bcolors.Colors(`(📜In the wilderness shall waters break out, & st.(Is.35:6)` + bcolors.ENDC)
    case 311:
        bcolors.Colors(`(📜He shall feed his flock like a shepherd.......(Is.40:11)` + bcolors.ENDC)
    case 312:
        bcolors.Colors(`(📜He giveth power to the faint; & to them that..(Is.40:29)` + bcolors.ENDC)
    case 313:
        bcolors.Colors(`(📜They that wait upon the LORD shall renew thei.(Is.40:31)` + bcolors.ENDC)
    case 314:
        bcolors.Colors(`(📜They shall run, & not be weary; & they shall..(Is.40:31)` + bcolors.ENDC)
    case 315:
        bcolors.Colors(`(📜I the LORD thy God will hold thy right hand...(Is.41:13)` + bcolors.ENDC)
    case 316:
        bcolors.Colors(`(📜Fear not; I will help thee....................(Is.41:13)` + bcolors.ENDC)
    case 317:
        bcolors.Colors(`(📜I will make the wilderness a pool of water....(Is.41:18)` + bcolors.ENDC)
    case 318:
        bcolors.Colors(`(📜I will make darkness light before them........(Is.42:16)` + bcolors.ENDC)
    case 319:
        bcolors.Colors(`(📜I will even make a way in the wilderness, & r.(Is.43:19)` + bcolors.ENDC)
    case 320:
        bcolors.Colors(`(📜I, even I, am he that blotteth out thy transg.(Is.43:25)` + bcolors.ENDC)
    case 321:
        bcolors.Colors(`(📜Return unto Me; for I have redeemed thee......(Is.44:22)` + bcolors.ENDC)
    case 322:
        bcolors.Colors(`(📜The word is gone out of my mouth in righteous.(Is.45:23)` + bcolors.ENDC)
    case 323:
        bcolors.Colors(`(📜I am the LORD thy God which teacheth thee to..(Is.48:17)` + bcolors.ENDC)
    case 324:
        bcolors.Colors(`(📜I have put my words in thy mouth..............(Is.51:16)` + bcolors.ENDC)
    case 325:
        bcolors.Colors(`(📜I have covered thee in the shadow of Mine.....(Is.51:16)` + bcolors.ENDC)
    case 326:
        bcolors.Colors(`(📜He is despised & rejected of men; a Man of.....(Is.53:3)` + bcolors.ENDC)
    case 327:
        bcolors.Colors(`(📜Surely he hath borne our griefs, & carried.....(Is.53:4)` + bcolors.ENDC)
    case 328:
        bcolors.Colors(`(📜He was wounded for our transgressions, He was..(Is.53:5)` + bcolors.ENDC)
    case 329:
        bcolors.Colors(`(📜The chastisement of our peace was upon Him.....(Is.53:5)` + bcolors.ENDC)
    case 330:
        bcolors.Colors(`(📜With his stripes we are healed.................(Is.53:5)` + bcolors.ENDC)
    case 331:
        bcolors.Colors(`(📜The LORD hath laid on him the iniquity of us...(Is.53:6)` + bcolors.ENDC)
    case 332:
        bcolors.Colors(`(📜He shall see of the travail of His soul, & sh.(Is.53:11)` + bcolors.ENDC)
    case 333:
        bcolors.Colors(`(📜He bare the sin of many, & made intercession..(Is.53:12)` + bcolors.ENDC)
    case 334:
        bcolors.Colors(`(📜Enlarge the place of thy tent..................(Is.54:2)` + bcolors.ENDC)
    case 335:
        bcolors.Colors(`(📜My kindness shall not depart from thee........(Is.54:10)` + bcolors.ENDC)
    case 336:
        bcolors.Colors(`(📜No weapon that is formed against thee shall...(Is.54:17)` + bcolors.ENDC)
    case 337:
        bcolors.Colors(`(📜Thou shalt be like a watered garden...........(Is.58:11)` + bcolors.ENDC)
    case 338:
        bcolors.Colors(`(📜The LORD shall be thine everlasting light.....(Is.60:20)` + bcolors.ENDC)
    case 339:
        bcolors.Colors(`(📜I all their affliction he was afflicted........(Is.63:9)` + bcolors.ENDC)
    case 340:
        bcolors.Colors(`(📜Ask for the old paths, where is the good way,.(Jer.6:16)` + bcolors.ENDC)
    case 341:
        bcolors.Colors(`(📜I am with thee to save thee & to deliver.....(Jer.15:20)` + bcolors.ENDC)
    case 342:
        bcolors.Colors(`(📜I the LORD search the heart, I try the reins (Jer.17:10)` + bcolors.ENDC)
    case 343:
        bcolors.Colors(`(📜Ye shall seek Me, & find Me, when ye shall...(Jer.29:13)` + bcolors.ENDC)
    case 344:
        bcolors.Colors(`(📜I will forgive their iniquity, & their sin I.(Jer.31:34)` + bcolors.ENDC)
    case 345:
        bcolors.Colors(`(📜I am the LORD, the God of all flesh: is ther.(Jer.32:27)` + bcolors.ENDC)
    case 346:
        bcolors.Colors(`(📜I will rejoice over them to do them good.....(Jer.32:41)` + bcolors.ENDC)
    case 347:
        bcolors.Colors(`(📜Call unto Me, & I will answer thee............(Jer.33:3)` + bcolors.ENDC)
    case 348:
        bcolors.Colors(`(📜I will shew thee great & mighty things, which.(Jer.33:3)` + bcolors.ENDC)
    case 349:
        bcolors.Colors(`(📜Obey, I beseech thee, the voice of the LORD..(Jer.38:20)` + bcolors.ENDC)
    case 350:
        bcolors.Colors(`(📜Come, & let us join ourselves to the LORD.....(Jer.50:5)` + bcolors.ENDC)
    case 351:
        bcolors.Colors(`(📜The LORD is good unto them that wait for him..(Lam.3:25)` + bcolors.ENDC)
    case 352:
        bcolors.Colors(`(📜Let us search & try our ways, & turn again....(Lam.3:40)` + bcolors.ENDC)
    case 353:
        bcolors.Colors(`(📜Let us lift up our heart with our hands unto..(Lam.3:41)` + bcolors.ENDC)
    case 354:
        bcolors.Colors(`(📜Cast away from you all your transgressions,.(Ezek.18:31)` + bcolors.ENDC)
    case 355:
        bcolors.Colors(`(📜I will accept you with your sweet savour....(Ezek.20:41)` + bcolors.ENDC)
    case 356:
        bcolors.Colors(`(📜I will cause the shower to come down in his.(Ezek.34:26)` + bcolors.ENDC)
    case 357:
        bcolors.Colors(`(📜There shall be showers of blessing..........(Ezek.34:26)` + bcolors.ENDC)
    case 358:
        bcolors.Colors(`(📜A new heart also will I give you, & a new...(Ezek.36:26)` + bcolors.ENDC)
    case 359:
        bcolors.Colors(`(📜Our God whom we serve is able to deliver......(Dan.3:17)` + bcolors.ENDC)
    case 360:
        bcolors.Colors(`(📜For I desired mercy, & not sacrifice...........(Hos.6:6)` + bcolors.ENDC)
    case 361:
        bcolors.Colors(`(📜Turn ye even to me with all your heart.......(Joel 2:12)` + bcolors.ENDC)
    case 362:
        bcolors.Colors(`(📜Be glad & rejoice in the LORD your God.......(Joel 2:23)` + bcolors.ENDC)
    case 363:
        bcolors.Colors(`(📜I will pour out My Spirit upon all flesh.....(Joel 2:28)` + bcolors.ENDC)
    case 364:
        bcolors.Colors(`(📜Whosoever shall call on the name of the LORD.(Joel 2:32)` + bcolors.ENDC)
    case 365:
        bcolors.Colors(`(📜The LORD will be the hope of his people......(Joel 3:16)` + bcolors.ENDC)
    case 366:
        bcolors.Colors(`(📜Seek ye me, & ye shall live...................(Amos 5:4)` + bcolors.ENDC)
    case 367:
        bcolors.Colors(`(📜The day of the LORD is near upon all the.....(Obad.1:15)` + bcolors.ENDC)
    case 368:
        bcolors.Colors(`(📜Salvation is of the LORD (Jon.2:9)` + bcolors.ENDC)
    case 369:
        bcolors.Colors(`(📜We will walk in the name of the LORD our God (Mic.4:5)` + bcolors.ENDC)
    case 370:
        bcolors.Colors(`(📜I will wait for the God of my salvation (Mic.7:7)` + bcolors.ENDC)
    case 371:
        bcolors.Colors(`(📜The LORD knoweth them that trust in him (Nah.1:7)` + bcolors.ENDC)
    case 372:
        bcolors.Colors(`(📜The just shall live by his faith (Hab.2:4)` + bcolors.ENDC)
    case 373:
        bcolors.Colors(`(📜O LORD, revive thy work in the midst of the years (Hab.3:2)` + bcolors.ENDC)
    case 374:
        bcolors.Colors(`(📜The LORD thy God in the midst of thee is mighty; He will save (Zeph.3:17)` + bcolors.ENDC)
    case 375:
        bcolors.Colors(`(📜Thus saith the LORD of hosts; Consider your ways. (Hag.1:7)` + bcolors.ENDC)
    case 376:
        bcolors.Colors(`(📜Be strong… for I am with you, saith the LORD (Hag.2:4)` + bcolors.ENDC)
    case 377:
        bcolors.Colors(`(📜My spirit remaineth among you: fear ye not! (Hag.2:5)` + bcolors.ENDC)
    case 378:
        bcolors.Colors(`(📜Fear not, but let your hands be strong (Zech.8:13)` + bcolors.ENDC)
    case 379:
        bcolors.Colors(`(📜For I am the LORD, I change not (Mal.3:6)` + bcolors.ENDC)
    case 380:
        bcolors.Colors(`(📜Bring ye all the tithes into the storehouse (Mal.3:10)` + bcolors.ENDC)
    case 381:
        bcolors.Colors(`(📜& all nations shall call you blessed (Mal.3:12)` + bcolors.ENDC)
    case 382:
        bcolors.Colors(`(📜Unto you that fear My name shall the Sun of Righteousness arise (Mal.4:2)` + bcolors.ENDC)
    case 383:
        bcolors.Colors(`(📜Prepare ye the way of the Lord, make His paths straight (Matt.3:3)` + bcolors.ENDC)
    case 384:
        bcolors.Colors(`(📜Therefore fruits meet for repentance (Matt.3:8)` + bcolors.ENDC)
    case 385:
        bcolors.Colors(`(📜He shall baptize you with the Holy Ghost, & with fire (Matt.3:11)` + bcolors.ENDC)
    case 386:
        bcolors.Colors(`(📜He will throughly purge His floor, & gather His wheat into the garner (Matt.3:12)` + bcolors.ENDC)
    case 387:
        bcolors.Colors(`(📜Man shall not live by bread alone (Matt.4:4)` + bcolors.ENDC)
    case 388:
        bcolors.Colors(`(📜Thou shalt worship the LORD thy God (Matt.4:10)` + bcolors.ENDC)
    case 389:
        bcolors.Colors(`(📜Repent: for the kingdom of heaven is at hand (Matt.4:17)` + bcolors.ENDC)
    case 390:
        bcolors.Colors(`(📜Follow Me, & I will make you fishers of men (Matt.4:19)` + bcolors.ENDC)
    case 391:
        bcolors.Colors(`(📜Blessed are the poor in spirit: for theirs is the kingdom of heaven (Matt.5:3)` + bcolors.ENDC)
    case 392:
        bcolors.Colors(`(📜Blessed are they that mourn: for they shall be comforted (Matt.5:4)` + bcolors.ENDC)
    case 393:
        bcolors.Colors(`(📜Blessed are the meek: for they shall inherit the earth (Matt.5:5)` + bcolors.ENDC)
    case 394:
        bcolors.Colors(`(📜Blessed are the merciful: for they shall obtain mercy (Matt.5:7)` + bcolors.ENDC)
    case 395:
        bcolors.Colors(`(📜Blessed are the pure in heart: for they shall see God. (Matt.5:8)` + bcolors.ENDC)
    case 396:
        bcolors.Colors(`(📜Blessed are the peacemakers: for they shall be called the children of God (Matt.5:9)` + bcolors.ENDC)
    case 397:
        bcolors.Colors(`(📜Ye are the light of the world. A city that is set on an hill cannot be hid. (Matt.5:14)` + bcolors.ENDC)
    case 398:
        bcolors.Colors(`(📜Let your light so shine before men (Matt.5:16)` + bcolors.ENDC)
    case 399:
        bcolors.Colors(`(📜Be ye therefore perfect, even as your Father which is in heaven is perfect. (Matt.5:48)` + bcolors.ENDC)
    case 400:
        bcolors.Colors(`(📜Your Father knoweth what things ye have need of (Matt.6:8)` + bcolors.ENDC)
    case 401:
        bcolors.Colors(`(📜But lay up for yourselves treasures in heaven (Matt.6:20)` + bcolors.ENDC)
    case 402:
        bcolors.Colors(`(📜For where your treasure is, there will your heart be also (Matt.6:21)` + bcolors.ENDC)
    case 403:
        bcolors.Colors(`(📜But seek ye first the kingdom of God, & his righteousness (Matt.6:33)` + bcolors.ENDC)
    case 404:
        bcolors.Colors(`(📜Ask, & it shall be given you (Matt.7:7)` + bcolors.ENDC)
    case 405:
        bcolors.Colors(`(📜Seek, & ye shall find (Matt.7:7)` + bcolors.ENDC)
    case 406:
        bcolors.Colors(`(📜Knock, & it shall be opened unto you (Matt.7:7)` + bcolors.ENDC)
    case 407:
        bcolors.Colors(`(📜Your Father which is in heaven give good things to them that ask him (Matt.7:11)` + bcolors.ENDC)
    case 408:
        bcolors.Colors(`(📜Narrow is the way, which leadeth unto life (Matt.7:14)` + bcolors.ENDC)
    case 409:
        bcolors.Colors(`(📜Be of good cheer; thy sins be forgiven thee (Matt.9:2)` + bcolors.ENDC)
    case 410:
        bcolors.Colors(`(📜Believe ye that I am able to do this? (Matt.9:28)` + bcolors.ENDC)
    case 411:
        bcolors.Colors(`(📜The harvest truly is plenteous, but the labourers are few (Matt.9:37)` + bcolors.ENDC)
    case 412:
        bcolors.Colors(`(📜Freely ye have received, freely give (Matt.10:8)` + bcolors.ENDC)
    case 413:
        bcolors.Colors(`(📜Be ye therefore wise as serpents, & harmless as doves (Matt.10:16)` + bcolors.ENDC)
    case 414:
        bcolors.Colors(`(📜What ye hear in the ear, that preach ye upon the housetops (Matt.10:27)` + bcolors.ENDC)
    case 415:
        bcolors.Colors(`(📜Come unto Me, all ye that labour & are heavy laden, & I will give you rest (Matt.11:28)` + bcolors.ENDC)
    case 416:
        bcolors.Colors(`(📜Learn of me; for I am meek & lowly in heart (Matt.11:29)` + bcolors.ENDC)
    case 417:
        bcolors.Colors(`(📜My yoke is easy, & My burden is light (Matt.11:30)` + bcolors.ENDC)
    case 418:
        bcolors.Colors(`(📜A bruised reed shall He not break (Matt.12:20)` + bcolors.ENDC)
    case 419:
        bcolors.Colors(`(📜He that gathereth not with Me scattereth abroad (Matt.12:30)` + bcolors.ENDC)
    case 420:
        bcolors.Colors(`(📜Then shall the righteous shine forth as the sun in the kingdom of their Father (Matt.13:43)` + bcolors.ENDC)
    case 421:
        bcolors.Colors(`(📜The kingdom of heaven is like unto treasure hid in a field (Matt.13:44)` + bcolors.ENDC)
    case 422:
        bcolors.Colors(`(📜I will build My church; & the gates of hell shall not prevail against it (Matt.16:18)` + bcolors.ENDC)
    case 423:
        bcolors.Colors(`(📜Many that are first shall be last; & the last shall be first (Matt.19:30)` + bcolors.ENDC)
    case 424:
        bcolors.Colors(`(📜He that shall endure unto the end, the same shall be saved (Matt.24:13)` + bcolors.ENDC)
    case 425:
        bcolors.Colors(`(📜& this gospel of the kingdom shall be preached in all the world (Matt.24:14)` + bcolors.ENDC)
    case 426:
        bcolors.Colors(`(📜Heaven & earth shall pass away, but My words shall not pass away (Matt.24:35)` + bcolors.ENDC)
    case 427:
        bcolors.Colors(`(📜Who then is a faithful & wise servant (Matt.24:45)` + bcolors.ENDC)
    case 428:
        bcolors.Colors(`(📜Watch & pray, that ye enter not into temptation (Matt.26:41)` + bcolors.ENDC)
    case 429:
        bcolors.Colors(`(📜The spirit indeed is willing, but the flesh is weak (Matt.26:41)` + bcolors.ENDC)
    case 430:
        bcolors.Colors(`(📜Go ye therefore, & teach all nations (Matt.28:19)` + bcolors.ENDC)
    case 431:
        bcolors.Colors(`(📜&, lo, I am with you alway, even unto the end of the world (Matt.28:20)` + bcolors.ENDC)
    case 432:
        bcolors.Colors(`(📜& in the morning, rising up a great while before day… prayed (Mark 1:35)` + bcolors.ENDC)
    case 433:
        bcolors.Colors(`(📜& as many as touched Him were made whole (Mark 6:56)` + bcolors.ENDC)
    case 434:
        bcolors.Colors(`(📜What would ye that I should do for you? (Mark 10:36)` + bcolors.ENDC)
    case 435:
        bcolors.Colors(`(📜Repent ye, & believe the gospel (Mark 11:15)` + bcolors.ENDC)
    case 436:
        bcolors.Colors(`(📜Go ye into all the world, & preach the gospel to every creature (Mark 16:15)` + bcolors.ENDC)
    case 437:
        bcolors.Colors(`(📜Blessed are ye that weep now: for ye shall laugh (Luke 6:21)` + bcolors.ENDC)
    case 438:
        bcolors.Colors(`(📜Take up his cross daily, & follow me (Luke 9:23)` + bcolors.ENDC)
    case 439:
        bcolors.Colors(`(📜In the beginning was the Word, & the Word was with God, & the Word was God (John 1:1)` + bcolors.ENDC)
    case 440:
        bcolors.Colors(`(📜& the light shineth in darkness; & the darkness comprehended it not. (John 1:5)` + bcolors.ENDC)
    case 441:
        bcolors.Colors(`(📜But as many as received him, to them gave he power to become the sons of God (John 1:12)` + bcolors.ENDC)
    case 442:
        bcolors.Colors(`(📜& the Word was made flesh, & dwelt among us (John 1:14)` + bcolors.ENDC)
    case 443:
        bcolors.Colors(`(📜Make straight the way of the LORD (John 1:23)` + bcolors.ENDC)
    case 444:
        bcolors.Colors(`(📜Hereafter ye shall see heaven open (John 1:51)` + bcolors.ENDC)
    case 445:
        bcolors.Colors(`(📜Except a man be born again, he cannot see the kingdom of God (John 3:3)` + bcolors.ENDC)
    case 446:
        bcolors.Colors(`(📜That which is born of the Spirit is spirit (John 3:6)` + bcolors.ENDC)
    case 447:
        bcolors.Colors(`(📜The wind bloweth where it listeth, & thou hearest the sound thereof (John 3:8)` + bcolors.ENDC)
    case 448:
        bcolors.Colors(`(📜For God so loved the world, that he gave his only begotten Son (John 3:16)` + bcolors.ENDC)
    case 449:
        bcolors.Colors(`(📜He that doeth truth cometh to the light (John 3:21)` + bcolors.ENDC)
    case 450:
        bcolors.Colors(`(📜He that believeth on the Son hath everlasting life (John 3:36)` + bcolors.ENDC)
    case 451:
        bcolors.Colors(`(📜The true worshippers shall worship the Father in spirit & in truth (John 4:23)` + bcolors.ENDC)
    case 452:
        bcolors.Colors(`(📜God is a Spirit: & they that worship Him must worship him in spirit & in truth (John 4:24)` + bcolors.ENDC)
    case 453:
        bcolors.Colors(`(📜Jesus saith unto them, My meat is to do the will of Him that sent Me, & to finish His work. (John 4:34)` + bcolors.ENDC)
    case 454:
        bcolors.Colors(`(📜He that reapeth receiveth wages, & gathereth fruit unto life eternal (John 4:36)` + bcolors.ENDC)
    case 456:
        bcolors.Colors(`(📜He that soweth & he that reapeth may rejoice together (John 4:36)` + bcolors.ENDC)
    case 457:
        bcolors.Colors(`(📜Behold, thou art made whole: sin no more (John 5:14)` + bcolors.ENDC)
    case 458:
        bcolors.Colors(`(📜Search the Scriptures; for in them ye think ye have eternal life (John 5:39)` + bcolors.ENDC)
    case 459:
        bcolors.Colors(`(📜Labour not for the meat which perisheth, but for that meat which endureth unto everlasting life (John 6:27)` + bcolors.ENDC)
    case 460:
        bcolors.Colors(`(📜This is the work of God, that ye believe on Him whom He hath sent (John 6:29)` + bcolors.ENDC)
    case 461:
        bcolors.Colors(`(📜For the bread of God is He which cometh down from heaven, & giveth life unto the world (John 6:33)` + bcolors.ENDC)
    case 462:
        bcolors.Colors(`(📜He that believeth on Me shall never thirst (John 6:35)` + bcolors.ENDC)
    case 463:
        bcolors.Colors(`(📜He that cometh to me shall never hunger (John 6:35)` + bcolors.ENDC)
    case 464:
        bcolors.Colors(`(📜Him that cometh to Me I will in no wise cast out (John 6:37)` + bcolors.ENDC)
    case 465:
        bcolors.Colors(`(📜He that believeth on Me hath everlasting life (John 6:47)` + bcolors.ENDC)
    case 466:
        bcolors.Colors(`(📜It is the spirit that quickeneth; the flesh profiteth nothing (John 6:63)` + bcolors.ENDC)
    case 467:
        bcolors.Colors(`(📜The words that I speak unto you, they are spirit, & they are life (John 6:63)` + bcolors.ENDC)
    case 468:
        bcolors.Colors(`(📜If any man thirst, let him come unto me, & drink (John 7:37)` + bcolors.ENDC)
    case 469:
        bcolors.Colors(`(📜He that believeth on Me… out of his belly shall flow rivers of living water (John 7:38)` + bcolors.ENDC)
    case 470:
        bcolors.Colors(`(📜He that followeth Me shall not walk in darkness (John 8:12)` + bcolors.ENDC)
    case 471:
        bcolors.Colors(`(📜If ye continue in My word, then are ye My disciples indeed (John 8:31)` + bcolors.ENDC)
    case 472:
        bcolors.Colors(`(📜& ye shall know the truth, & the truth shall make you free (John 8:32)` + bcolors.ENDC)
    case 473:
        bcolors.Colors(`(📜Whosoever committeth sin is the servant of sin (John 8:34)` + bcolors.ENDC)
    case 474:
        bcolors.Colors(`(📜He that is of God heareth God's words (John 8:47)` + bcolors.ENDC)
    case 475:
        bcolors.Colors(`(📜If a man keep My saying, he shall never see death (John 8:51)` + bcolors.ENDC)
    case 476:
        bcolors.Colors(`(📜I am the door: by Me if any man enter in, he shall be saved. (John 10:9)` + bcolors.ENDC)
    case 477:
        bcolors.Colors(`(📜I am come that they might have life, & that they might have it more abundantly (John 10:10)` + bcolors.ENDC)
    case 478:
        bcolors.Colors(`(📜I am the good shepherd: the good shepherd giveth His life for the sheep (John 10:11)` + bcolors.ENDC)
    case 479:
        bcolors.Colors(`(📜I am the good shepherd, & know My sheep, & am known of Mine (John 10:14)` + bcolors.ENDC)
    case 480:
        bcolors.Colors(`(📜My sheep hear My voice, & I know them, & they follow Me. (John 10:27)` + bcolors.ENDC)
    case 481:
        bcolors.Colors(`(📜Neither shall any man pluck them out of My hand (John 10:28)` + bcolors.ENDC)
    case 482:
        bcolors.Colors(`(📜I am the resurrection, & the life (John 11:25)` + bcolors.ENDC)
    case 483:
        bcolors.Colors(`(📜If thou wouldest believe, thou shouldest see the glory of God (John 11:40)` + bcolors.ENDC)
    case 484:
        bcolors.Colors(`(📜If any man serve Me, him will My Father honour (John 12:26)` + bcolors.ENDC)
    case 485:
        bcolors.Colors(`(📜If any man serve me, let him follow Me (John 12:26)` + bcolors.ENDC)
    case 486:
        bcolors.Colors(`(📜While ye have light, believe in the light, that ye may be the children of light (John 12:36)` + bcolors.ENDC)
    case 487:
        bcolors.Colors(`(📜A new comm&ment I give unto you, that ye love one another (John 13:34)` + bcolors.ENDC)
    case 488:
        bcolors.Colors(`(📜By this shall all men know that ye are My disciples, if ye have love one to another (John 13:35)` + bcolors.ENDC)
    case 489:
        bcolors.Colors(`(📜Let not your heart be troubled: ye believe in God, believe also in Me. (John 14:1)` + bcolors.ENDC)
    case 490:
        bcolors.Colors(`(📜I will come again, & receive you unto Myself (John 14:3)` + bcolors.ENDC)
    case 491:
        bcolors.Colors(`(📜I am the way, the truth, & the life (John 14:6)` + bcolors.ENDC)
    case 492:
        bcolors.Colors(`(📜He that believeth on Me, the works that I do shall he do also (John 14:12)` + bcolors.ENDC)
    case 493:
        bcolors.Colors(`(📜& whatsoever ye shall ask in My name, that will I do (John 14:13)` + bcolors.ENDC)
    case 494:
        bcolors.Colors(`(📜If ye love Me, keep My commandments (John 14:15)` + bcolors.ENDC)
    case 495:
        bcolors.Colors(`(📜The Comforter, which is the Holy Ghost… He shall teach you all things (John 14:26)` + bcolors.ENDC)
    case 496:
        bcolors.Colors(`(📜Let not your heart be troubled, neither let it be afraid (John 14:27)` + bcolors.ENDC)
    case 497:
        bcolors.Colors(`(📜Abide in Me, & I in you (John 15:4)` + bcolors.ENDC)
    case 498:
        bcolors.Colors(`(📜He that abideth in Me, & I in him, the same bringeth forth much fruit (John 15:5)` + bcolors.ENDC)
    case 499:
        bcolors.Colors(`(📜Herein is My Father glorified, that ye bear much fruit (John 15:8)` + bcolors.ENDC)
    case 500:
        bcolors.Colors(`(📜I loved you: continue ye in My love (John 15:9)` + bcolors.ENDC)
    case 501:
        bcolors.Colors(`(📜That My joy might remain in you, & that your joy might be full. (John 15:11)` + bcolors.ENDC)
    case 502:
        bcolors.Colors(`(📜This is My comm&ment, that ye love one another, as I have loved you (John 15:12)` + bcolors.ENDC)
    case 503:
        bcolors.Colors(`(📜Ye are My friends, if ye do whatsoever I comm& you. (John 15:14)` + bcolors.ENDC)
    case 504:
        bcolors.Colors(`(📜Ye have not chosen Me, but I have chosen you (John 15:16)` + bcolors.ENDC)
    case 505:
        bcolors.Colors(`(📜I have chosen you, & ordained you, that ye should go & bring forth fruit (John 15:16)` + bcolors.ENDC)
    case 506:
        bcolors.Colors(`(📜Ask, & ye shall receive, that your joy may be full (John 16:24)` + bcolors.ENDC)
    case 507:
        bcolors.Colors(`(📜Be of good cheer; I have overcome the world (John 16:33)` + bcolors.ENDC)
    case 508:
        bcolors.Colors(`(📜As my Father hath sent Me, even so send I you (John 20:21)` + bcolors.ENDC)
    case 509:
        bcolors.Colors(`(📜Blessed are they that have not seen, & yet have believed (John 20:29)` + bcolors.ENDC)
    case 510:
        bcolors.Colors(`(📜Lord,… grant unto Thy servants, that with all boldness they may speak Thy word (Acts 4:29)` + bcolors.ENDC)
    case 511:
        bcolors.Colors(`(📜The will of the Lord be done (Acts 21:14)` + bcolors.ENDC)
    case 512:
        bcolors.Colors(`(📜I am not ashamed of the gospel of Christ (Rom.1:16)` + bcolors.ENDC)
    case 513:
        bcolors.Colors(`(📜The just shall live by faith (Rom.1:17)` + bcolors.ENDC)
    case 514:
        bcolors.Colors(`(📜The goodness of God leadeth thee to repentance (Rom.2:4)` + bcolors.ENDC)
    case 515:
        bcolors.Colors(`(📜All have sinned, & come short of the glory of God (Rom.3:23)` + bcolors.ENDC)
    case 516:
        bcolors.Colors(`(📜Blessed are they whose iniquities are forgiven (Rom.4:7)` + bcolors.ENDC)
    case 517:
        bcolors.Colors(`(📜Being justified by faith, we have peace with God (Rom.5:1)` + bcolors.ENDC)
    case 518:
        bcolors.Colors(`(📜The love of God is shed abroad in our hearts by the Holy Ghost (Rom.5:5)` + bcolors.ENDC)
    case 519:
        bcolors.Colors(`(📜While we were yet sinners, Christ died for us (Rom.5:8)` + bcolors.ENDC)
    case 520:
        bcolors.Colors(`(📜Being now justified by His blood, we shall be saved from wrath through Him (Rom.5:9)` + bcolors.ENDC)
    case 521:
        bcolors.Colors(`(📜When we were enemies, we were reconciled to God by the death of His Son (Rom.5:10)` + bcolors.ENDC)
    case 522:
        bcolors.Colors(`(📜Where sin abounded, grace did much more abound (Rom.5:20)` + bcolors.ENDC)
    case 523:
        bcolors.Colors(`(📜Yield yourselves unto God, as those that are alive from the dead (Rom.6:13)` + bcolors.ENDC)
    case 524:
        bcolors.Colors(`(📜The gift of God is eternal life through Jesus Christ our Lord (Rom.6:23)` + bcolors.ENDC)
    case 525:
        bcolors.Colors(`(📜We should bring forth fruit unto God (Rom.7:4)` + bcolors.ENDC)
    case 526:
        bcolors.Colors(`(📜We are debtors, not to the flesh, to live after the flesh (Rom.8:12)` + bcolors.ENDC)
    case 527:
        bcolors.Colors(`(📜For as many as are led by the Spirit of God, they are the sons of God. (Rom.8:14)` + bcolors.ENDC)
    case 528:
        bcolors.Colors(`(📜Likewise the Spirit also helpeth our infirmities. (Rom.8:26)` + bcolors.ENDC)
    case 529:
        bcolors.Colors(`(📜All things work together for good to them that love God (Rom.8:28)` + bcolors.ENDC)
    case 530:
        bcolors.Colors(`(📜If God be for us, who can be against us? (Rom.8:31)` + bcolors.ENDC)
    case 531:
        bcolors.Colors(`(📜Who shall lay any thing to the charge of God's elect? It is God that justifieth. (Rom.8:33)` + bcolors.ENDC)
    case 532:
        bcolors.Colors(`(📜Nay, in all these things we are more than conquerors through him that loved us. (Rom.8:37)` + bcolors.ENDC)
    case 533:
        bcolors.Colors(`(📜Nor height, nor depth, nor any other creature, shall be able to separate us from the love of God (Rom.8:39)` + bcolors.ENDC)
    case 534:
        bcolors.Colors(`(📜O man, who art thou that repliest against God? (Rom.9:20)` + bcolors.ENDC)
    case 535:
        bcolors.Colors(`(📜Whosoever believeth on him shall not be ashamed (Rom.9:33)` + bcolors.ENDC)
    case 536:
        bcolors.Colors(`(📜The word is nigh thee, even in thy mouth, & in thy heart (Rom.10:8)` + bcolors.ENDC)
    case 537:
        bcolors.Colors(`(📜With the mouth confession is made unto salvation (Rom.10:10)` + bcolors.ENDC)
    case 538:
        bcolors.Colors(`(📜The same Lord over all is rich unto all that call upon Him. (Rom.10:12)` + bcolors.ENDC)
    case 539:
        bcolors.Colors(`(📜How beautiful are the feet of them that preach the gospel of peace (Rom.10:15)` + bcolors.ENDC)
    case 540:
        bcolors.Colors(`(📜Faith cometh by hearing, & hearing by the word of God (Rom.10:17)` + bcolors.ENDC)
    case 541:
        bcolors.Colors(`(📜The gifts & calling of God are without repentance (Rom.11:29)` + bcolors.ENDC)
    case 542:
        bcolors.Colors(`(📜O the depth of the riches both of the wisdom & knowledge of God! (Rom.11:33)` + bcolors.ENDC)
    case 543:
        bcolors.Colors(`(📜& be not conformed to this world (Rom.12:2)` + bcolors.ENDC)
    case 544:
        bcolors.Colors(`(📜Transformed by the renewing of your mind (Rom.12:2)` + bcolors.ENDC)
    case 545:
        bcolors.Colors(`(📜That ye may prove what is that good, & acceptable, & perfect, will of God (Rom.12:2)` + bcolors.ENDC)
    case 546:
        bcolors.Colors(`(📜Let love be without dissimulation (Rom.12:9)` + bcolors.ENDC)
    case 547:
        bcolors.Colors(`(📜Be kindly affectioned one to another with brotherly love (Rom.12:10)` + bcolors.ENDC)
    case 548:
        bcolors.Colors(`(📜Fervent in spirit; serving the Lord (Rom.12:11)` + bcolors.ENDC)
    case 549:
        bcolors.Colors(`(📜Rejoicing in hope; patient in tribulation; continuing instant in prayer (Rom.12:12)` + bcolors.ENDC)
    case 550:
        bcolors.Colors(`(📜Rejoice with them that do rejoice, & weep with them that weep (Rom.12:15)` + bcolors.ENDC)
    case 551:
        bcolors.Colors(`(📜Recompense to no man evil for evil (Rom.12:17)` + bcolors.ENDC)
    case 552:
        bcolors.Colors(`(📜Live peaceably with all men (Rom.12:18)` + bcolors.ENDC)
    case 553:
        bcolors.Colors(`(📜Be not overcome of evil, but overcome evil with good (Rom.12:21)` + bcolors.ENDC)
    case 554:
        bcolors.Colors(`(📜Owe no man any thing, but to love one another (Rom.13:8)` + bcolors.ENDC)
    case 555:
        bcolors.Colors(`(📜Now it is high time to awake out of sleep (Rom.13:11)` + bcolors.ENDC)
    case 556:
        bcolors.Colors(`(📜Now is our salvation nearer than when we believed (Rom.13:11)` + bcolors.ENDC)
    case 557:
        bcolors.Colors(`(📜Let us put on the armour of light (Rom.13:12)` + bcolors.ENDC)
    case 558:
        bcolors.Colors(`(📜Put ye on the Lord Jesus Christ (Rom.13:14)` + bcolors.ENDC)
    case 559:
        bcolors.Colors(`(📜So then every one of us shall give account of himself to God (Rom.14:12)` + bcolors.ENDC)
    case 560:
        bcolors.Colors(`(📜The kingdom of God is… righteousness, & peace, & joy in the Holy Ghost. (Rom.14:17)` + bcolors.ENDC)
    case 561:
        bcolors.Colors(`(📜Happy is he that condemneth not himself in that thing which he alloweth (Rom.14:22)` + bcolors.ENDC)
    case 562:
        bcolors.Colors(`(📜Christ… received us to the glory of God (Rom.15:7)` + bcolors.ENDC)
    case 563:
        bcolors.Colors(`(📜Now the God of hope fill you with all joy & peace in believing (Rom.15:13)` + bcolors.ENDC)
    case 564:
        bcolors.Colors(`(📜The God of peace shall bruise Satan under your feet shortly (Rom.16:20)` + bcolors.ENDC)
    case 565:
        bcolors.Colors(`(📜For the preaching of the cross is… the power of God (1Cor.1:18)` + bcolors.ENDC)
    case 567:
        bcolors.Colors(`(📜We preach Christ crucified (1Cor.1:23)` + bcolors.ENDC)
    case 568:
        bcolors.Colors(`(📜God hath chosen the weak things of the world to confound the things which are mighty (1Cor.1:27)` + bcolors.ENDC)
    case 569:
        bcolors.Colors(`(📜But God hath chosen the foolish things of the world to confound the wise (1Cor.1:27)` + bcolors.ENDC)
    case 570:
        bcolors.Colors(`(📜We have received, not the spirit of the world, but the spirit which is of God (1Cor.2:12)` + bcolors.ENDC)
    case 571:
        bcolors.Colors(`(📜Every man shall receive his own reward according to his own labour (1Cor.3:8)` + bcolors.ENDC)
    case 572:
        bcolors.Colors(`(📜We are labourers together with God (1Cor.3:9)` + bcolors.ENDC)
    case 573:
        bcolors.Colors(`(📜Know ye not that ye are the temple of God? (1Cor.3:16)` + bcolors.ENDC)
    case 574:
        bcolors.Colors(`(📜The temple of God is holy, which temple ye are (1Cor.3:17)` + bcolors.ENDC)
    case 575:
        bcolors.Colors(`(📜Ye are Christ's; & Christ is God's (1Cor.3:23)` + bcolors.ENDC)
    case 576:
        bcolors.Colors(`(📜We are fools for Christ's sake (1Cor.4:10)` + bcolors.ENDC)
    case 577:
        bcolors.Colors(`(📜The kingdom of God is not in word, but in power (1Cor.4:20)` + bcolors.ENDC)
    case 578:
        bcolors.Colors(`(📜For even Christ our passover is sacrificed for us (1Cor.5:7)` + bcolors.ENDC)
    case 579:
        bcolors.Colors(`(📜He that is joined unto the Lord is one spirit (1Cor.6:17)` + bcolors.ENDC)
    case 580:
        bcolors.Colors(`(📜Ye are bought with a price (1Cor.6:20)` + bcolors.ENDC)
    case 581:
        bcolors.Colors(`(📜Glorify God in your body, & in your spirit, which are God's (1Cor.6:20)` + bcolors.ENDC)
    case 582:
        bcolors.Colors(`(📜For the fashion of this world passeth away (1Cor.7:31)` + bcolors.ENDC)
    case 583:
        bcolors.Colors(`(📜Knowledge puffeth up, but charity edifieth (1Cor.8:1)` + bcolors.ENDC)
    case 584:
        bcolors.Colors(`(📜If any man love God, the same is known of Him (1Cor.8:3)` + bcolors.ENDC)
    case 585:
        bcolors.Colors(`(📜So run, that ye may obtain (1Cor.9:24)` + bcolors.ENDC)
    case 586:
        bcolors.Colors(`(📜Let him that thinketh he st&eth take heed lest he fall (1Cor.10:12)` + bcolors.ENDC)
    case 587:
        bcolors.Colors(`(📜God is faithful, who will not suffer you to be tempted above that ye are able (1Cor.10:13)` + bcolors.ENDC)
    case 589:
        bcolors.Colors(`(📜Do all to the glory of God (1Cor.10:31)` + bcolors.ENDC)
    case 590:
        bcolors.Colors(`(📜The manifestation of the Spirit is given to every man (1Cor.12:7)` + bcolors.ENDC)
    case 591:
        bcolors.Colors(`(📜Covet earnestly the best gifts (1Cor.12:31)` + bcolors.ENDC)
    case 592:
        bcolors.Colors(`(📜Charity suffereth long, & is kind (1Cor.13:4)` + bcolors.ENDC)
    case 593:
        bcolors.Colors(`(📜Charity vaunteth not itself, is not puffed up (1Cor.13:4)` + bcolors.ENDC)
    case 594:
        bcolors.Colors(`(📜Charity never faileth (1Cor.13:8)` + bcolors.ENDC)
    case 595:
        bcolors.Colors(`(📜& now abideth faith, hope, charity, these three (1Cor.13:13)` + bcolors.ENDC)
    case 596:
        bcolors.Colors(`(📜Follow after charity, & desire spiritual gifts (1Cor.14:1)` + bcolors.ENDC)
    case 597:
        bcolors.Colors(`(📜Christ died for our sins according to the Scriptures (1Cor.15:3)` + bcolors.ENDC)
    case 598:
        bcolors.Colors(`(📜But now is Christ risen from the dead (1Cor.15:20)` + bcolors.ENDC)
    case 599:
        bcolors.Colors(`(📜Evil communications corrupt good manners (1Cor.15:33)` + bcolors.ENDC)
    case 600:
        bcolors.Colors(`(📜Thanks be to God, which giveth us the victory (1Cor.15:57)` + bcolors.ENDC)
    case 601:
        bcolors.Colors(`(📜Be ye stedfast, unmoveable (1Cor.15:58)` + bcolors.ENDC)
    case 602:
        bcolors.Colors(`(📜Always abounding in the work of the Lord (1Cor.15:58)` + bcolors.ENDC)
    case 603:
        bcolors.Colors(`(📜Your labour is not in vain in the Lord (1Cor.15:58)` + bcolors.ENDC)
    case 604:
        bcolors.Colors(`(📜Watch ye, st& fast in the faith, quit you like men, be strong (1Cor.16:13)` + bcolors.ENDC)
    case 605:
        bcolors.Colors(`(📜Let all your things be done with charity (1Cor.16:14)` + bcolors.ENDC)
    case 606:
        bcolors.Colors(`(📜All the promises of God in Him are yea, & in Him Amen (2Cor.1:20)` + bcolors.ENDC)
    case 607:
        bcolors.Colors(`(📜He which stablisheth us with you in Christ, & hath anointed us, is God (2Cor.1:21)` + bcolors.ENDC)
    case 608:
        bcolors.Colors(`(📜Thanks be unto God, which always causeth us to triumph in Christ (2Cor.2:14)` + bcolors.ENDC)
    case 609:
        bcolors.Colors(`(📜We are unto God a sweet savour of Christ (2Cor.2:15)` + bcolors.ENDC)
    case 610:
        bcolors.Colors(`(📜Ye are… the epistle of Christ (2Cor.3:3)` + bcolors.ENDC)
    case 611:
        bcolors.Colors(`(📜Our sufficiency is of God (2Cor.3:5)` + bcolors.ENDC)
    case 612:
        bcolors.Colors(`(📜Where the Spirit of the Lord is, there is liberty (2Cor.3:17)` + bcolors.ENDC)
    case 613:
        bcolors.Colors(`(📜God, who comm&ed the light to shine out of darkness, hath shined in our hearts (2Cor.4:6)` + bcolors.ENDC)
    case 614:
        bcolors.Colors(`(📜We have this treasure in earthen vessels (2Cor.4:7)` + bcolors.ENDC)
    case 615:
        bcolors.Colors(`(📜We are troubled on every side, yet not distressed (2Cor.4:8)` + bcolors.ENDC)
    case 616:
        bcolors.Colors(`(📜We are perplexed, but not in despair (2Cor.4:8)` + bcolors.ENDC)
    case 617:
        bcolors.Colors(`(📜Our light affliction… worketh for us a far more exceeding & eternal weight of glory (2Cor.4:17)` + bcolors.ENDC)
    case 618:
        bcolors.Colors(`(📜The things which are seen are temporal; but the things which are not seen are eternal (2Cor.4:18)` + bcolors.ENDC)
    case 619:
        bcolors.Colors(`(📜We have a building of God… in the heavens (2Cor.5:1)` + bcolors.ENDC)
    case 620:
        bcolors.Colors(`(📜We walk by faith, not by sight (2Cor.5:7)` + bcolors.ENDC)
    case 621:
        bcolors.Colors(`(📜If any man be in Christ, he is a new creature (2Cor.5:17)` + bcolors.ENDC)
    case 622:
        bcolors.Colors(`(📜Old things are passed away; behold, all things are become new (2Cor.5:17)` + bcolors.ENDC)
    case 623:
        bcolors.Colors(`(📜Be ye reconciled to God (2Cor.5:20)` + bcolors.ENDC)
    case 624:
        bcolors.Colors(`(📜He hath made him to be sin for us, who knew no sin (2Cor.5:21)` + bcolors.ENDC)
    case 625:
        bcolors.Colors(`(📜Behold, now is the accepted time; behold, now is the day of salvation (2Cor.6:2)` + bcolors.ENDC)
    case 626:
        bcolors.Colors(`(📜As sorrowful, yet alway rejoicing (2Cor.6:10)` + bcolors.ENDC)
    case 627:
        bcolors.Colors(`(📜Touch not the unclean thing; & I will receive you (2Cor.6:17)` + bcolors.ENDC)
    case 628:
        bcolors.Colors(`(📜& will be a Father unto you, & ye shall be my sons & daughters (2Cor.6:18)` + bcolors.ENDC)
    case 629:
        bcolors.Colors(`(📜Having therefore these promises, dearly beloved, let us cleanse ourselves (2Cor.7:1)` + bcolors.ENDC)
    case 630:
        bcolors.Colors(`(📜For godly sorrow worketh repentance to salvation (2Cor.7:10)` + bcolors.ENDC)
    case 631:
        bcolors.Colors(`(📜Yet for your sakes He became poor, that ye through His poverty might be rich (2Cor.8:9)` + bcolors.ENDC)
    case 632:
        bcolors.Colors(`(📜Providing for honest things, not only in the sight of the Lord, but also in the sight of men (2Cor.8:21)` + bcolors.ENDC)
    case 633:
        bcolors.Colors(`(📜He which soweth sparingly shall reap also sparingly (2Cor.9:6)` + bcolors.ENDC)
    case 634:
        bcolors.Colors(`(📜God loveth a cheerful giver (2Cor.9:7)` + bcolors.ENDC)
    case 635:
        bcolors.Colors(`(📜God is able to make all grace abound toward you (2Cor.9:8)` + bcolors.ENDC)
    case 636:
        bcolors.Colors(`(📜That ye… may abound to every good work (2Cor.9:8)` + bcolors.ENDC)
    case 637:
        bcolors.Colors(`(📜He that ministereth seed to the sower… multiply your seed sown (2Cor.9:10)` + bcolors.ENDC)
    case 638:
        bcolors.Colors(`(📜Thanks be unto God for his unspeakable gift (2Cor.9:15)` + bcolors.ENDC)
    case 639:
        bcolors.Colors(`(📜For though we walk in the flesh, we do not war after the flesh. (2Cor.10:3)` + bcolors.ENDC)
    case 640:
        bcolors.Colors(`(📜The weapons of our warfare are not carnal (2Cor.10:4)` + bcolors.ENDC)
    case 641:
        bcolors.Colors(`(📜Bringing into captivity every thought to the obedience of Christ (2Cor.10:5)` + bcolors.ENDC)
    case 642:
        bcolors.Colors(`(📜But he that glorieth, let him glory in the LORD (2Cor.10:17)` + bcolors.ENDC)
    case 643:
        bcolors.Colors(`(📜My strength is made perfect in weakness (2Cor.12:9)` + bcolors.ENDC)
    case 644:
        bcolors.Colors(`(📜Be perfect, be of good comfort, be of one mind, live in peace. (2Cor.13:11)` + bcolors.ENDC)
    case 645:
        bcolors.Colors(`(📜Be of one mind, live in peace (2Cor.13:11)` + bcolors.ENDC)
    case 646:
        bcolors.Colors(`(📜Who gave Himself for our sins (Gal.1:4)` + bcolors.ENDC)
    case 647:
        bcolors.Colors(`(📜God accepteth no man's person (Gal.2:6)` + bcolors.ENDC)
    case 648:
        bcolors.Colors(`(📜We have believed in Jesus Christ, that we might be justified by the faith of Christ (Gal.2:16)` + bcolors.ENDC)
    case 649:
        bcolors.Colors(`(📜Nevertheless I live; yet not I, but Christ liveth in me (Gal.2:20)` + bcolors.ENDC)
    case 650:
        bcolors.Colors(`(📜They which are of faith, the same are the children of Abraham (Gal.3:7)` + bcolors.ENDC)
    case 651:
        bcolors.Colors(`(📜Christ hath redeemed us from the curse of the law (Gal.3:13)` + bcolors.ENDC)
    case 652:
        bcolors.Colors(`(📜If ye be Christ's, then are ye Abraham's seed, & heirs according to the promise (Gal.3:29)` + bcolors.ENDC)
    case 653:
        bcolors.Colors(`(📜God hath sent forth the Spirit of his Son into your hearts (Gal.4:6)` + bcolors.ENDC)
    case 654:
        bcolors.Colors(`(📜Wherefore thou art no more a servant, but a son; & if a son, then an heir of God (Gal.4:7)` + bcolors.ENDC)
    case 655:
        bcolors.Colors(`(📜St& fast therefore in the liberty wherewith Christ hath made us free (Gal.5:1)` + bcolors.ENDC)
    case 656:
        bcolors.Colors(`(📜By love serve one another (Gal.5:13)` + bcolors.ENDC)
    case 657:
        bcolors.Colors(`(📜Thou shalt love thy neighbour as thyself (Gal.5:14)` + bcolors.ENDC)
    case 658:
        bcolors.Colors(`(📜Walk in the Spirit (Gal.5:16)` + bcolors.ENDC)
    case 659:
        bcolors.Colors(`(📜The fruit of the Spirit is love, joy, peace… (Gal.5:22)` + bcolors.ENDC)
    case 660:
        bcolors.Colors(`(📜They that are Christ's have crucified the flesh (Gal.5:24)` + bcolors.ENDC)
    case 661:
        bcolors.Colors(`(📜If we live in the Spirit, let us also walk in the Spirit (Gal.5:25)` + bcolors.ENDC)
    case 662:
        bcolors.Colors(`(📜Bear ye one another's burdens, & so fulfil the law of Christ. (Gal.6:2)` + bcolors.ENDC)
    case 663:
        bcolors.Colors(`(📜Let every man prove his own work (Gal.6:4)` + bcolors.ENDC)
    case 664:
        bcolors.Colors(`(📜Whatsoever a man soweth, that shall he also reap (Gal.6:7)` + bcolors.ENDC)
    case 665:
        bcolors.Colors(`(📜He that soweth to the Spirit shall of the Spirit reap life everlasting (Gal.6:8)` + bcolors.ENDC)
    case 666:
        bcolors.Colors(`(📜Let us not be weary in well doing (Gal.6:9)` + bcolors.ENDC)
    case 667:
        bcolors.Colors(`(📜As we have therefore opportunity, let us do good unto all men (Gal.6:10)` + bcolors.ENDC)
    case 668:
        bcolors.Colors(`(📜Blessed be the God,… who hath blessed us with every spiritual blessing (Eph.1:3)` + bcolors.ENDC)
    case 669:
        bcolors.Colors(`(📜He hath chosen us in Him before the foundation of the world (Eph.1:4)` + bcolors.ENDC)
    case 670:
        bcolors.Colors(`(📜In whom we have redemption through His blood, the forgiveness of sins (Eph.1:7)` + bcolors.ENDC)
    case 671:
        bcolors.Colors(`(📜Ye were sealed with that holy Spirit of promise (Eph.1:13)` + bcolors.ENDC)
    case 672:
        bcolors.Colors(`(📜What the riches of the glory of His inheritance in the saints (Eph.1:18)` + bcolors.ENDC)
    case 673:
        bcolors.Colors(`(📜What is the exceeding greatness of His power to us (Eph.1:19)` + bcolors.ENDC)
    case 674:
        bcolors.Colors(`(📜For by grace are ye saved through faith (Eph.2:8)` + bcolors.ENDC)
    case 675:
        bcolors.Colors(`(📜For we are his workmanship, created in Christ Jesus unto good works (Eph.2:10)` + bcolors.ENDC)
    case 676:
        bcolors.Colors(`(📜Ye who sometimes were far off are made nigh by the blood of Christ (Eph.2:13)` + bcolors.ENDC)
    case 677:
        bcolors.Colors(`(📜Ye are… fellowcitizens with the saints, & of the household of God (Eph.2:19)` + bcolors.ENDC)
    case 678:
        bcolors.Colors(`(📜I beseech you that ye walk worthy of the vocation wherewith ye are called (Eph.4:1)` + bcolors.ENDC)
    case 679:
        bcolors.Colors(`(📜Unto every one of us is given grace according to the measure of the gift of Christ. (Eph.4:7)` + bcolors.ENDC)
    case 680:
        bcolors.Colors(`(📜Speak every man truth with his neighbour (Eph.4:25)` + bcolors.ENDC)
    case 681:
        bcolors.Colors(`(📜Grieve not the holy Spirit of God (Eph.4:30)` + bcolors.ENDC)
    case 682:
        bcolors.Colors(`(📜Be ye kind one to another, tenderhearted (Eph.4:32)` + bcolors.ENDC)
    case 683:
        bcolors.Colors(`(📜Forgiving one another, even as God for Christ's sake hath forgiven you (Eph.4:32)` + bcolors.ENDC)
    case 684:
        bcolors.Colors(`(📜Be ye therefore followers of God, as dear children (Eph.5:1)` + bcolors.ENDC)
    case 685:
        bcolors.Colors(`(📜Christ also hath loved us, & hath given Himself for us (Eph.5:2)` + bcolors.ENDC)
    case 686:
        bcolors.Colors(`(📜Walk as children of light (Eph.5:8)` + bcolors.ENDC)
    case 687:
        bcolors.Colors(`(📜The fruit of the Spirit is in all goodness & righteousness & truth (Eph.5:9)` + bcolors.ENDC)
    case 688:
        bcolors.Colors(`(📜Proving what is acceptable unto the Lord (Eph.5:10)` + bcolors.ENDC)
    case 689:
        bcolors.Colors(`(📜See then that ye walk circumspectly, not as fools, but as wise (Eph.5:15)` + bcolors.ENDC)
    case 690:
        bcolors.Colors(`(📜Redeeming the time, because the days are evil (Eph.5:16)` + bcolors.ENDC)
    case 691:
        bcolors.Colors(`(📜Underst&ing what the will of the Lord is (Eph.5:17)` + bcolors.ENDC)
    case 692:
        bcolors.Colors(`(📜Be filled with the Spirit (Eph.5:18)` + bcolors.ENDC)
    case 693:
        bcolors.Colors(`(📜Christ also loved the church, & gave Himself for it (Eph.5:25)` + bcolors.ENDC)
    case 694:
        bcolors.Colors(`(📜Be strong in the Lord, & in the power of his might (Eph.6:10)` + bcolors.ENDC)
    case 695:
        bcolors.Colors(`(📜Put on the whole armour of God (Eph.6:11)` + bcolors.ENDC)
    case 696:
        bcolors.Colors(`(📜We wrestle not against flesh & blood (Eph.6:12)` + bcolors.ENDC)
    case 697:
        bcolors.Colors(`(📜St& therefore, having your loins girt about with truth (Eph.6:14)` + bcolors.ENDC)
    case 698:
        bcolors.Colors(`(📜Praying always with all prayer & supplication in the Spirit (Eph.6:18)` + bcolors.ENDC)
    case 699:
        bcolors.Colors(`(📜For to me to live is Christ, & to die is gain (Phil.1:21)` + bcolors.ENDC)
    case 700:
        bcolors.Colors(`(📜Only let your conversation be as it becometh the gospel of Christ (Phil.1:27)` + bcolors.ENDC)
    case 701:
        bcolors.Colors(`(📜Let nothing be done through strife or vainglory (Phil.2:3)` + bcolors.ENDC)
    case 702:
        bcolors.Colors(`(📜In lowliness of mind let each esteem other better than themselves (Phil.2:3)` + bcolors.ENDC)
    case 703:
        bcolors.Colors(`(📜Work out your own salvation with fear & trembling (Phil.2:12)` + bcolors.ENDC)
    case 704:
        bcolors.Colors(`(📜It is God which worketh in you both to will & to do of his good pleasure (Phil.2:13)` + bcolors.ENDC)
    case 705:
        bcolors.Colors(`(📜Do all things without murmurings & disputings (Phil.2:14)` + bcolors.ENDC)
    case 706:
        bcolors.Colors(`(📜I count all things but loss for the excellency of the knowledge of Christ Jesus (Phil.3:8)` + bcolors.ENDC)
    case 707:
        bcolors.Colors(`(📜Forgetting those things which are behind, & reaching forth unto those things which are before (Phil.3:13)` + bcolors.ENDC)
    case 708:
        bcolors.Colors(`(📜Our conversation is in heaven (Phil.3:20)` + bcolors.ENDC)
    case 709:
        bcolors.Colors(`(📜Rejoice in the Lord alway: & again I say, Rejoice! (Phil.4:4)` + bcolors.ENDC)
    case 710:
        bcolors.Colors(`(📜Let your moderation be known unto all men (Phil.4:5)` + bcolors.ENDC)
    case 711:
        bcolors.Colors(`(📜The Lord is at hand (Phil.4:5)` + bcolors.ENDC)
    case 712:
        bcolors.Colors(`(📜Be careful for nothing (Phil.4:6)` + bcolors.ENDC)
    case 713:
        bcolors.Colors(`(📜Let your requests be made known unto God (Phil.4:6)` + bcolors.ENDC)
    case 714:
        bcolors.Colors(`(📜I am instructed both to be full & to be hungry, both to abound & to suffer need (Phil.4:12)` + bcolors.ENDC)
    case 715:
        bcolors.Colors(`(📜I can do all things through Christ which strengtheneth me (Phil.4:13)` + bcolors.ENDC)
    case 716:
        bcolors.Colors(`(📜God shall supply all your need (Phil.4:19)` + bcolors.ENDC)
    case 717:
        bcolors.Colors(`(📜That ye might be filled with the knowledge of His will (Col.1:9)` + bcolors.ENDC)
    case 718:
        bcolors.Colors(`(📜That ye might walk worthy of the Lord (Col.1:10)` + bcolors.ENDC)
    case 719:
        bcolors.Colors(`(📜He hath delivered us from the power of darkness (Col.1:13)` + bcolors.ENDC)
    case 720:
        bcolors.Colors(`(📜In whom we have redemption through His blood, even the forgiveness of sins (Col.1:14)` + bcolors.ENDC)
    case 721:
        bcolors.Colors(`(📜Christ in you, the hope of glory (Col.1:27)` + bcolors.ENDC)
    case 722:
        bcolors.Colors(`(📜As ye have therefore received Christ Jesus the Lord, so walk ye in Him (Col.2:6)` + bcolors.ENDC)
    case 723:
        bcolors.Colors(`(📜In Him dwelleth all the fulness of the Godhead bodily (Col.2:9)` + bcolors.ENDC)
    case 724:
        bcolors.Colors(`(📜Ye are complete in Him (Col.2:10)` + bcolors.ENDC)
    case 725:
        bcolors.Colors(`(📜Hath he quickened together with Him, having forgiven you all trespasses (Col.2:13)` + bcolors.ENDC)
    case 726:
        bcolors.Colors(`(📜Seek those things which are above, where Christ sitteth on the right hand of God (Col.3:1)` + bcolors.ENDC)
    case 727:
        bcolors.Colors(`(📜Set your affection on things above, not on things on the earth (Col.3:2)` + bcolors.ENDC)
    case 728:
        bcolors.Colors(`(📜Your life is hid with Christ in God (Col.3:3)` + bcolors.ENDC)
    case 729:
        bcolors.Colors(`(📜& above all these things put on charity (Col.3:14)` + bcolors.ENDC)
    case 730:
        bcolors.Colors(`(📜Charity… is the bond of perfectness (Col.3:14)` + bcolors.ENDC)
    case 731:
        bcolors.Colors(`(📜Let the peace of God rule in your hearts (Col.3:15)` + bcolors.ENDC)
    case 732:
        bcolors.Colors(`(📜Be ye thankful (Col.3:15)` + bcolors.ENDC)
    case 733:
        bcolors.Colors(`(📜Let the word of Christ dwell in you richly (Col.3:16)` + bcolors.ENDC)
    case 734:
        bcolors.Colors(`(📜Do all in the name of the Lord Jesus (Col.3:17)` + bcolors.ENDC)
    case 735:
        bcolors.Colors(`(📜& whatsoever ye do, do it heartily, as to the Lord (Col.3:23)` + bcolors.ENDC)
    case 736:
        bcolors.Colors(`(📜Of the Lord ye shall receive the reward of the inheritance (Col.3:24)` + bcolors.ENDC)
    case 737:
        bcolors.Colors(`(📜Continue in prayer (Col.4:2)` + bcolors.ENDC)
    case 738:
        bcolors.Colors(`(📜Let your speech be alway with grace, seasoned with salt (Col.4:6)` + bcolors.ENDC)
    case 739:
        bcolors.Colors(`(📜Take heed to the ministry which thou hast received in the Lord, that thou fulfil it (Col.4:17)` + bcolors.ENDC)
    case 740:
        bcolors.Colors(`(📜& the Lord make you to increase & abound in love one toward another (1Thess.3:12)` + bcolors.ENDC)
    case 741:
        bcolors.Colors(`(📜For this is the will of God, even your sanctification (1Thess.4:3)` + bcolors.ENDC)
    case 742:
        bcolors.Colors(`(📜That every one of you should know how to possess his vessel in sanctification & honour (1Thess.4:4)` + bcolors.ENDC)
    case 743:
        bcolors.Colors(`(📜For God hath not called us unto uncleanness, but unto holiness (1Thess.4:7)` + bcolors.ENDC)
    case 744:
        bcolors.Colors(`(📜That ye may walk honestly toward them that are without (1Thess.4:12)` + bcolors.ENDC)
    case 745:
        bcolors.Colors(`(📜Ye are all the children of light, & the children of the day (1Thess.5:5)` + bcolors.ENDC)
    case 746:
        bcolors.Colors(`(📜Let us not sleep, as do others; but let us watch & be sober (1Thess.5:6)` + bcolors.ENDC)
    case 748:
        bcolors.Colors(`(📜Let us, who are of the day, be sober (1Thess.5:8)` + bcolors.ENDC)
    case 749:
        bcolors.Colors(`(📜For God hath not appointed us to wrath, but to obtain salvation (1Thess.5:9)` + bcolors.ENDC)
    case 750:
        bcolors.Colors(`(📜Ever follow that which is good (1Thess.5:15)` + bcolors.ENDC)
    case 751:
        bcolors.Colors(`(📜Rejoice evermore (1Thess.5:16)` + bcolors.ENDC)
    case 752:
        bcolors.Colors(`(📜Pray without ceasing (1Thess.5:17)` + bcolors.ENDC)
    case 753:
        bcolors.Colors(`(📜In every thing give thanks (1Thess.5:18)` + bcolors.ENDC)
    case 754:
        bcolors.Colors(`(📜Prove all things; hold fast that which is good (1Thess.5:21)` + bcolors.ENDC)
    case 755:
        bcolors.Colors(`(📜The very God of peace sanctify you wholly (1Thess.5:23)` + bcolors.ENDC)
    case 756:
        bcolors.Colors(`(📜I pray God your whole spirit & soul & body be preserved blameless (1Thess.5:23)` + bcolors.ENDC)
    case 757:
        bcolors.Colors(`(📜God hath from the beginning chosen you to salvation (2Thess.2:13)` + bcolors.ENDC)
    case 758:
        bcolors.Colors(`(📜The Lord is faithful, who shall stablish you, & keep you from evil (2Thess.3:3)` + bcolors.ENDC)
    case 759:
        bcolors.Colors(`(📜If any would not work, neither should he eat (2Thess.3:10)` + bcolors.ENDC)
    case 760:
        bcolors.Colors(`(📜The Lord of peace himself give you peace always by all means (2Thess.3:16)` + bcolors.ENDC)
    case 761:
        bcolors.Colors(`(📜I thank Christ Jesus our Lord, who hath enabled me (1Tim.1:12)` + bcolors.ENDC)
    case 762:
        bcolors.Colors(`(📜Christ Jesus came into the world to save sinners (1Tim.1:15)` + bcolors.ENDC)
    case 763:
        bcolors.Colors(`(📜There is one God, & one mediator between God & men (1Tim.2:5)` + bcolors.ENDC)
    case 764:
        bcolors.Colors(`(📜Great is the mystery of godliness: God was manifest in the flesh (1Tim.3:16)` + bcolors.ENDC)
    case 765:
        bcolors.Colors(`(📜We trust in the living God, who is the Saviour of all men (1Tim.4:10)` + bcolors.ENDC)
    case 766:
        bcolors.Colors(`(📜Till I come, give attendance to reading, to exhortation, to doctrine. (1Tim.4:13)` + bcolors.ENDC)
    case 767:
        bcolors.Colors(`(📜Neglect not the gift that is in thee, which was given thee (1Tim.4:14)` + bcolors.ENDC)
    case 768:
        bcolors.Colors(`(📜Take heed unto thyself, & unto the doctrine; continue in them. (1Tim.4:16)` + bcolors.ENDC)
    case 769:
        bcolors.Colors(`(📜Thou shalt not muzzle the ox that treadeth out the corn (1Tim.5:18)` + bcolors.ENDC)
    case 770:
        bcolors.Colors(`(📜The labourer is worthy of his reward (1Tim.5:18)` + bcolors.ENDC)
    case 771:
        bcolors.Colors(`(📜Keep thyself pure (1Tim.5:22)` + bcolors.ENDC)
    case 772:
        bcolors.Colors(`(📜Godliness with contentment is great gain (1Tim.6:6)` + bcolors.ENDC)
    case 773:
        bcolors.Colors(`(📜Having food & raiment let us be therewith content (1Tim.6:8)` + bcolors.ENDC)
    case 774:
        bcolors.Colors(`(📜The love of money is the root of all evil (1Tim.6:10)` + bcolors.ENDC)
    case 775:
        bcolors.Colors(`(📜Follow after righteousness, godliness, faith, love, patience, meekness (1Tim.6:11)` + bcolors.ENDC)
    case 776:
        bcolors.Colors(`(📜Fight the good fight of faith (1Tim.6:12)` + bcolors.ENDC)
    case 777:
        bcolors.Colors(`(📜Lay hold on eternal life, whereunto thou art also called (1Tim.6:12)` + bcolors.ENDC)
    case 778:
        bcolors.Colors(`(📜Keep that which is committed to thy trust (1Tim.6:20)` + bcolors.ENDC)
    case 779:
        bcolors.Colors(`(📜I put thee in remembrance that thou stir up the gift of God, which is in thee (2Tim.1:6)` + bcolors.ENDC)
    case 780:
        bcolors.Colors(`(📜God hath not given us the spirit of fear; but of power, & of love, & of a sound mind (2Tim.1:7)` + bcolors.ENDC)
    case 781:
        bcolors.Colors(`(📜Be not thou therefore ashamed of the testimony of our Lord (2Tim.1:8)` + bcolors.ENDC)
    case 782:
        bcolors.Colors(`(📜Hold fast the form of sound words (2Tim.1:13)` + bcolors.ENDC)
    case 783:
        bcolors.Colors(`(📜That good thing which was committed unto thee keep by the Holy Ghost, who dwelleth in us. (2Tim.1:14)` + bcolors.ENDC)
    case 784:
        bcolors.Colors(`(📜Be strong in the grace that is in Christ Jesus (2Tim.2:1)` + bcolors.ENDC)
    case 785:
        bcolors.Colors(`(📜Thou therefore endure hardness, as a good soldier of Jesus Christ (2Tim.2:3)` + bcolors.ENDC)
    case 786:
        bcolors.Colors(`(📜The husb&man that laboureth must be first partaker of the fruits. (2Tim.2:6)` + bcolors.ENDC)
    case 787:
        bcolors.Colors(`(📜Remember that Jesus Christ of the seed of David was raised from the dead (2Tim.2:8)` + bcolors.ENDC)
    case 788:
        bcolors.Colors(`(📜If we suffer, we shall also reign with Him (2Tim.2:12)` + bcolors.ENDC)
    case 789:
        bcolors.Colors(`(📜The Lord knoweth them that are His (2Tim.2:19)` + bcolors.ENDC)
    case 790:
        bcolors.Colors(`(📜Follow righteousness, faith, charity, peace, with them that call on the Lord (2Tim.2:22)` + bcolors.ENDC)
    case 791:
        bcolors.Colors(`(📜The servant of the Lord must not strive; but be gentle unto all men (2Tim.2:24)` + bcolors.ENDC)
    case 792:
        bcolors.Colors(`(📜All that will live godly in Christ Jesus shall suffer persecution (2Tim.3:12)` + bcolors.ENDC)
    case 793:
        bcolors.Colors(`(📜Continue thou in the things which thou hast learned (2Tim.3:14)` + bcolors.ENDC)
    case 794:
        bcolors.Colors(`(📜All Scripture is given by inspiration of God, & is profitable for doctrine (2Tim.3:16)` + bcolors.ENDC)
    case 795:
        bcolors.Colors(`(📜That the man of God may be perfect, throughly furnished unto all good works (2Tim.3:17)` + bcolors.ENDC)
    case 796:
        bcolors.Colors(`(📜Preach the word; be instant in season, out of season. (2Tim.4:2)` + bcolors.ENDC)
    case 797:
        bcolors.Colors(`(📜But watch thou in all things (2Tim.4:5)` + bcolors.ENDC)
    case 798:
        bcolors.Colors(`(📜Do the work of an evangelist, make full proof of thy ministry. (2Tim.4:5)` + bcolors.ENDC)
    case 799:
        bcolors.Colors(`(📜In all things shewing thyself a pattern of good works (Titus 2:7)` + bcolors.ENDC)
    case 800:
        bcolors.Colors(`(📜The grace of God that bringeth salvation hath appeared to all men (Titus 2:11)` + bcolors.ENDC)
    case 801:
        bcolors.Colors(`(📜Not by works of righteousness which we have done, but according to His mercy He saved us (Titus 3:5)` + bcolors.ENDC)
    case 802:
        bcolors.Colors(`(📜That the communication of thy faith may become effectual (Philem.1:6)` + bcolors.ENDC)
    case 803:
        bcolors.Colors(`(📜Having confidence in thy obedience I wrote unto thee (Philem.1:21)` + bcolors.ENDC)
    case 804:
        bcolors.Colors(`(📜We ought to give the more earnest heed to the things which we have heard (Heb.2:1)` + bcolors.ENDC)
    case 805:
        bcolors.Colors(`(📜For both He that sanctifieth & they who are sanctified are all of one (Heb.2:11)` + bcolors.ENDC)
    case 806:
        bcolors.Colors(`(📜As the children are partakers of flesh & blood, he also himself likewise took part of the same (Heb.2:14)` + bcolors.ENDC)
    case 807:
        bcolors.Colors(`(📜He is able to succour them that are tempted (Heb.2:18)` + bcolors.ENDC)
    case 808:
        bcolors.Colors(`(📜For every house is builded by some man; but he that built all things is God. (Heb.3:4)` + bcolors.ENDC)
    case 809:
        bcolors.Colors(`(📜Christ as a son over His own house; whose house are we (Heb.3:6)` + bcolors.ENDC)
    case 810:
        bcolors.Colors(`(📜Exhort one another daily (Heb.3:13)` + bcolors.ENDC)
    case 811:
        bcolors.Colors(`(📜For we are made partakers of Christ (Heb.3:14)` + bcolors.ENDC)
    case 812:
        bcolors.Colors(`(📜We which have believed do enter into rest (Heb.4:3)` + bcolors.ENDC)
    case 813:
        bcolors.Colors(`(📜Let us labour therefore to enter into that rest (Heb.4:11)` + bcolors.ENDC)
    case 814:
        bcolors.Colors(`(📜The word of God is quick, & powerful, & sharper than any twoedged sword (Heb.4:12)` + bcolors.ENDC)
    case 815:
        bcolors.Colors(`(📜All things are naked & opened unto the eyes of Him with whom we have to do. (Heb.4:13)` + bcolors.ENDC)
    case 816:
        bcolors.Colors(`(📜Let us hold fast our profession (Heb.4:14)` + bcolors.ENDC)
    case 817:
        bcolors.Colors(`(📜Let us therefore come boldly unto the throne of grace (Heb.4:16)` + bcolors.ENDC)
    case 818:
        bcolors.Colors(`(📜Yet learned He obedience by the things which He suffered (Heb.5:8)` + bcolors.ENDC)
    case 819:
        bcolors.Colors(`(📜Leaving the principles of the doctrine of Christ, let us go on unto perfection (Heb.6:1)` + bcolors.ENDC)
    case 820:
        bcolors.Colors(`(📜For God is not unrighteous to forget your work & labour of love (Heb.6:10)` + bcolors.ENDC)
    case 821:
        bcolors.Colors(`(📜Surely blessing I will bless thee, & multiplying I will multiply thee (Heb.6:14)` + bcolors.ENDC)
    case 822:
        bcolors.Colors(`(📜Jesus… is able also to save them to the uttermost that come unto God by Him (Heb.7:25)` + bcolors.ENDC)
    case 823:
        bcolors.Colors(`(📜Jesus… ever liveth to make intercession (Heb.7:25)` + bcolors.ENDC)
    case 824:
        bcolors.Colors(`(📜I will be to them a God, & they shall be to me a people (Heb.8:10)` + bcolors.ENDC)
    case 825:
        bcolors.Colors(`(📜I will put my laws into their mind, & write them in their hearts (Heb.8:10)` + bcolors.ENDC)
    case 826:
        bcolors.Colors(`(📜The blood of Christ… purge your conscience from dead works (Heb.9:14)` + bcolors.ENDC)
    case 827:
        bcolors.Colors(`(📜We are sanctified through the offering of the body of Jesus Christ (Heb.10:10)` + bcolors.ENDC)
    case 828:
        bcolors.Colors(`(📜Their sins & iniquities will I remember no more (Heb.10:17)` + bcolors.ENDC)
    case 829:
        bcolors.Colors(`(📜Let us hold fast the profession of our faith without wavering (Heb.10:23)` + bcolors.ENDC)
    case 830:
        bcolors.Colors(`(📜Let us consider one another to provoke unto love & to good works (Heb.10:24)` + bcolors.ENDC)
    case 831:
        bcolors.Colors(`(📜Not forsaking the assembling of ourselves together (Heb.10:25)` + bcolors.ENDC)
    case 832:
        bcolors.Colors(`(📜Cast not away therefore your confidence, which hath great recompence of reward (Heb.10:35)` + bcolors.ENDC)
    case 833:
        bcolors.Colors(`(📜Ye have need of patience, that, after ye have done the will of God, ye might receive the promise (Heb.10:36)` + bcolors.ENDC)
    case 834:
        bcolors.Colors(`(📜Yet a little while, & he that shall come will come, & will not tarry (Heb.10:37)` + bcolors.ENDC)
    case 835:
        bcolors.Colors(`(📜We are… of them that believe to the saving of the soul (Heb.10:39)` + bcolors.ENDC)
    case 836:
        bcolors.Colors(`(📜Faith is the substance of things hoped for, the evidence of things not seen. (Heb.11:1)` + bcolors.ENDC)
    case 837:
        bcolors.Colors(`(📜But without faith it is impossible to please Him (Heb.11:6)` + bcolors.ENDC)
    case 838:
        bcolors.Colors(`(📜He that cometh to God must believe that He is (Heb.11:6)` + bcolors.ENDC)
    case 839:
        bcolors.Colors(`(📜God is a rewarder of them that diligently seek him (Heb.11:6)` + bcolors.ENDC)
    case 840:
        bcolors.Colors(`(📜Strangers & pilgrims on the earth (Heb.11:13)` + bcolors.ENDC)
    case 841:
        bcolors.Colors(`(📜Through faith subdued kingdoms, wrought righteousness, obtained promises (Heb.11:33)` + bcolors.ENDC)
    case 842:
        bcolors.Colors(`(📜God having provided some better thing for us (Heb.11:40)` + bcolors.ENDC)
    case 843:
        bcolors.Colors(`(📜Let us lay aside every weight, & the sin which doth so easily beset us (Heb.12:1)` + bcolors.ENDC)
    case 844:
        bcolors.Colors(`(📜Let us run with patience the race that is set before us (Heb.12:1)` + bcolors.ENDC)
    case 845:
        bcolors.Colors(`(📜For the joy that was set before Him endured the cross (Heb.12:2)` + bcolors.ENDC)
    case 846:
        bcolors.Colors(`(📜Consider Him that endured such contradiction of sinners against Himself (Heb.12:3)` + bcolors.ENDC)
    case 847:
        bcolors.Colors(`(📜Ye have not yet resisted unto blood, striving against sin (Heb.12:4)` + bcolors.ENDC)
    case 848:
        bcolors.Colors(`(📜Nor faint when thou art rebuked of Him (Heb.12:5)` + bcolors.ENDC)
    case 849:
        bcolors.Colors(`(📜Lift up the hands which hang down, & the feeble knees (Heb.12:12)` + bcolors.ENDC)
    case 850:
        bcolors.Colors(`(📜Follow peace with all men, & holiness (Heb.12:14)` + bcolors.ENDC)
    case 851:
        bcolors.Colors(`(📜Ye are come unto mount Sion, & unto the city of the living God (Heb.12:22)` + bcolors.ENDC)
    case 852:
        bcolors.Colors(`(📜Wherefore we receiving a kingdom which cannot be moved, let us have grace (Heb.12:28)` + bcolors.ENDC)
    case 853:
        bcolors.Colors(`(📜Let us have grace, whereby we may serve God acceptably with reverence & godly fear (Heb.12:28)` + bcolors.ENDC)
    case 854:
        bcolors.Colors(`(📜Let brotherly love continue. (Heb.13:1)` + bcolors.ENDC)
    case 855:
        bcolors.Colors(`(📜Be not forgetful to entertain strangers (Heb.13:2)` + bcolors.ENDC)
    case 856:
        bcolors.Colors(`(📜Be content with such things as ye have (Heb.13:5)` + bcolors.ENDC)
    case 857:
        bcolors.Colors(`(📜I will never leave thee, nor forsake thee (Heb.13:5)` + bcolors.ENDC)
    case 858:
        bcolors.Colors(`(📜The LORD is my helper, & I will not fear what man shall do unto me (Heb.13:6)` + bcolors.ENDC)
    case 859:
        bcolors.Colors(`(📜Jesus Christ the same yesterday, & to day, & for ever (Heb.13:8)` + bcolors.ENDC)
    case 860:
        bcolors.Colors(`(📜It is a good thing that the heart be established with grace (Heb.13:9)` + bcolors.ENDC)
    case 861:
        bcolors.Colors(`(📜Jesus, that he might sanctify the people with His own blood, suffered without the gate (Heb.13:12)` + bcolors.ENDC)
    case 862:
        bcolors.Colors(`(📜Let us go forth therefore unto Him without the camp, bearing His reproach. (Heb.13:13)` + bcolors.ENDC)
    case 863:
        bcolors.Colors(`(📜Here have we no continuing city, but we seek one to come (Heb.13:14)` + bcolors.ENDC)
    case 864:
        bcolors.Colors(`(📜By Him let us offer the sacrifice of praise to God continually, that is, the fruit of our lips (Heb.13:15)` + bcolors.ENDC)
    case 865:
        bcolors.Colors(`(📜To do good & to communicate forget not (Heb.13:16)` + bcolors.ENDC)
    case 866:
        bcolors.Colors(`(📜Obey them that have the rule over you (Heb.13:17)` + bcolors.ENDC)
    case 867:
        bcolors.Colors(`(📜He that wavereth is like a wave of the sea (James 1:6)` + bcolors.ENDC)
    case 868:
        bcolors.Colors(`(📜Blessed is the man that endureth temptation (James 1:12)` + bcolors.ENDC)
    case 869:
        bcolors.Colors(`(📜Every good gift & every perfect gift is from above (James 1:17)` + bcolors.ENDC)
    case 870:
        bcolors.Colors(`(📜Of His own will begat he us with the word of truth (James 1:18)` + bcolors.ENDC)
    case 871:
        bcolors.Colors(`(📜Receive with meekness the engrafted word (James 1:21)` + bcolors.ENDC)
    case 872:
        bcolors.Colors(`(📜Be ye doers of the word (James 1:22)` + bcolors.ENDC)
    case 873:
        bcolors.Colors(`(📜To keep himself unspotted from the world (James 1:27)` + bcolors.ENDC)
    case 874:
        bcolors.Colors(`(📜Mercy rejoiceth against judgment. (James 2:13)` + bcolors.ENDC)
    case 875:
        bcolors.Colors(`(📜By works was faith made perfect (James 2:22)` + bcolors.ENDC)
    case 876:
        bcolors.Colors(`(📜Faith without works is dead (James 2:26)` + bcolors.ENDC)
    case 877:
        bcolors.Colors(`(📜If any man offend not in word, the same is a perfect man (James 3:2)` + bcolors.ENDC)
    case 878:
        bcolors.Colors(`(📜The spirit that dwelleth in us lusteth to envy (James 4:5)` + bcolors.ENDC)
    case 879:
        bcolors.Colors(`(📜God resisteth the proud, but giveth grace unto the humble (James 4:6)` + bcolors.ENDC)
    case 880:
        bcolors.Colors(`(📜Submit yourselves therefore to God. Resist the devil. (James 4:7)` + bcolors.ENDC)
    case 881:
        bcolors.Colors(`(📜Draw nigh to God, & He will draw nigh to you. (James 4:8)` + bcolors.ENDC)
    case 882:
        bcolors.Colors(`(📜Humble yourselves in the sight of the Lord, & He shall lift you up. (James 4:10)` + bcolors.ENDC)
    case 883:
        bcolors.Colors(`(📜Be patient therefore, brethren, unto the coming of the Lord (James 5:7)` + bcolors.ENDC)
    case 884:
        bcolors.Colors(`(📜Be ye also patient; stablish your hearts (James 5:8)` + bcolors.ENDC)
    case 885:
        bcolors.Colors(`(📜The coming of the Lord draweth nigh (James 5:8)` + bcolors.ENDC)
    case 886:
        bcolors.Colors(`(📜That the Lord is very pitiful, & of tender mercy (James 5:11)` + bcolors.ENDC)
    case 887:
        bcolors.Colors(`(📜Is any among you afflicted? let him pray. (James 5:13)` + bcolors.ENDC)
    case 888:
        bcolors.Colors(`(📜Is any merry? let him sing psalms. (James 5:13)` + bcolors.ENDC)
    case 889:
        bcolors.Colors(`(📜The prayer of faith shall save the sick (James 5:15)` + bcolors.ENDC)
    case 890:
        bcolors.Colors(`(📜Pray one for another, that ye may be healed (James 5:16)` + bcolors.ENDC)
    case 891:
        bcolors.Colors(`(📜The effectual fervent prayer of a righteous man availeth much (James 5:16)` + bcolors.ENDC)
    case 892:
        bcolors.Colors(`(📜Blessed be the God… which hath begotten us again unto a lively hope (1Pet.1:3)` + bcolors.ENDC)
    case 893:
        bcolors.Colors(`(📜Be ye holy in all manner of conversation (1Pet.1:15)` + bcolors.ENDC)
    case 894:
        bcolors.Colors(`(📜Be ye holy; for I am holy (1Pet.1:16)` + bcolors.ENDC)
    case 895:
        bcolors.Colors(`(📜Ye were not redeemed with corruptible things, as silver & gold, from your vain conversation (1Pet.1:18)` + bcolors.ENDC)
    case 896:
        bcolors.Colors(`(📜Love one another with a pure heart fervently (1Pet.1:22)` + bcolors.ENDC)
    case 897:
        bcolors.Colors(`(📜The word of the Lord endureth for ever (1Pet.1:25)` + bcolors.ENDC)
    case 898:
        bcolors.Colors(`(📜As newborn babes, desire the sincere milk of the word (1Pet.2:2)` + bcolors.ENDC)
    case 899:
        bcolors.Colors(`(📜Ye have tasted that the Lord is gracious (1Pet.2:3)` + bcolors.ENDC)
    case 900:
        bcolors.Colors(`(📜He that believeth on him shall not be confounded (1Pet.2:6)` + bcolors.ENDC)
    case 901:
        bcolors.Colors(`(📜But ye are a chosen generation, a royal priesthood, an holy nation (1Pet.2:9)` + bcolors.ENDC)
    case 902:
        bcolors.Colors(`(📜Honour all men. Love the brotherhood. Fear God. Honour the king. (1Pet.2:17)` + bcolors.ENDC)
    case 903:
        bcolors.Colors(`(📜Who His own self bare our sins in His own body on the tree (1Pet.2:24)` + bcolors.ENDC)
    case 904:
        bcolors.Colors(`(📜By whose stripes ye were healed (1Pet.2:24)` + bcolors.ENDC)
    case 905:
        bcolors.Colors(`(📜Ye were as sheep going astray; but are now returned unto the Shepherd (1Pet.2:25)` + bcolors.ENDC)
    case 906:
        bcolors.Colors(`(📜The eyes of the Lord are over the righteous (1Pet.3:12)` + bcolors.ENDC)
    case 907:
        bcolors.Colors(`(📜If ye suffer for righteousness' sake, happy are ye (1Pet.3:14)` + bcolors.ENDC)
    case 908:
        bcolors.Colors(`(📜Sanctify the Lord God in your hearts (1Pet.3:15)` + bcolors.ENDC)
    case 909:
        bcolors.Colors(`(📜Christ hath once suffered for sins,… that He might bring us to God (1Pet.3:18)` + bcolors.ENDC)
    case 910:
        bcolors.Colors(`(📜The end of all things is at hand (1Pet.4:7)` + bcolors.ENDC)
    case 911:
        bcolors.Colors(`(📜Use hospitality one to another without grudging (1Pet.4:9)` + bcolors.ENDC)
    case 912:
        bcolors.Colors(`(📜As every man hath received the gift, even so minister the same one to another (1Pet.4:10)` + bcolors.ENDC)
    case 913:
        bcolors.Colors(`(📜If any man speak, let him speak as the oracles of God (1Pet.4:11)` + bcolors.ENDC)
    case 914:
        bcolors.Colors(`(📜If any man minister, let him do it as of the ability which God giveth (1Pet.4:11)` + bcolors.ENDC)
    case 915:
        bcolors.Colors(`(📜The spirit of glory & of God resteth upon you (1Pet.4:14)` + bcolors.ENDC)
    case 916:
        bcolors.Colors(`(📜Ye shall receive a crown of glory that fadeth not away (1Pet.5:4)` + bcolors.ENDC)
    case 917:
        bcolors.Colors(`(📜Humble yourselves therefore under the mighty hand of God, that He may exalt you in due time (1Pet.5:6)` + bcolors.ENDC)
    case 918:
        bcolors.Colors(`(📜Casting all your care upon him; for he careth for you (1Pet.5:7)` + bcolors.ENDC)
    case 919:
        bcolors.Colors(`(📜As His divine power hath given unto us all things that pertain unto life & godliness (2Pet.1:3)` + bcolors.ENDC)
    case 920:
        bcolors.Colors(`(📜Exceeding great & precious promises (2Pet.1:4)` + bcolors.ENDC)
    case 921:
        bcolors.Colors(`(📜Give diligence to make your calling & election sure (2Pet.1:10)` + bcolors.ENDC)
    case 922:
        bcolors.Colors(`(📜The Lord knoweth how to deliver the godly out of temptations (2Pet.2:9)` + bcolors.ENDC)
    case 923:
        bcolors.Colors(`(📜The Lord is not slack concerning His promise (2Pet.3:9)` + bcolors.ENDC)
    case 924:
        bcolors.Colors(`(📜The day of the Lord will come as a thief in the night (2Pet.3:10)` + bcolors.ENDC)
    case 925:
        bcolors.Colors(`(📜Be diligent that ye may be found of him in peace, without spot, & blameless (2Pet.3:14)` + bcolors.ENDC)
    case 926:
        bcolors.Colors(`(📜Account that the longsuffering of our Lord is salvation (2Pet.3:15)` + bcolors.ENDC)
    case 927:
        bcolors.Colors(`(📜Grow in grace, & in the knowledge of our Lord (2Pet.3:18)` + bcolors.ENDC)
    case 928:
        bcolors.Colors(`(📜For the life was manifested (1John 1:2)` + bcolors.ENDC)
    case 929:
        bcolors.Colors(`(📜Our fellowship is with the Father, & with His Son Jesus Christ (1John 1:3)` + bcolors.ENDC)
    case 930:
        bcolors.Colors(`(📜That your joy may be full (1John 1:4)` + bcolors.ENDC)
    case 931:
        bcolors.Colors(`(📜God is light, & in him is no darkness at all (1John 1:5)` + bcolors.ENDC)
    case 932:
        bcolors.Colors(`(📜The blood of Jesus Christ his Son cleanseth us from all sin (1John 1:7)` + bcolors.ENDC)
    case 933:
        bcolors.Colors(`(📜We have an advocate with the Father, Jesus Christ the righteous (1John 2:1)` + bcolors.ENDC)
    case 934:
        bcolors.Colors(`(📜Whoso keepeth His word, in him verily is the love of God perfected (1John 2:5)` + bcolors.ENDC)
    case 935:
        bcolors.Colors(`(📜The darkness is past, & the true light now shineth (1John 2:8)` + bcolors.ENDC)
    case 936:
        bcolors.Colors(`(📜Your sins are forgiven you for His name's sake (1John 2:12)` + bcolors.ENDC)
    case 937:
        bcolors.Colors(`(📜& the world passeth away, & the lust thereof (1John 2:17)` + bcolors.ENDC)
    case 938:
        bcolors.Colors(`(📜He that doeth the will of God abideth for ever (1John 2:17)` + bcolors.ENDC)
    case 939:
        bcolors.Colors(`(📜This is the promise that he hath promised us, even eternal life (1John 2:25)` + bcolors.ENDC)
    case 940:
        bcolors.Colors(`(📜& now, little children, abide in Him (1John 2:28)` + bcolors.ENDC)
    case 941:
        bcolors.Colors(`(📜Behold, what manner of love the Father hath bestowed upon us (1John 3:1)` + bcolors.ENDC)
    case 942:
        bcolors.Colors(`(📜Beloved, now are we the sons of God (1John 3:2)` + bcolors.ENDC)
    case 943:
        bcolors.Colors(`(📜He was manifested to take away our sins (1John 3:5)` + bcolors.ENDC)
    case 944:
        bcolors.Colors(`(📜Hereby perceive we the love of God, because He laid down His life for us (1John 3:16)` + bcolors.ENDC)
    case 945:
        bcolors.Colors(`(📜Let us not love in word, neither in tongue; but in deed & in truth (1John 3:18)` + bcolors.ENDC)
    case 946:
        bcolors.Colors(`(📜If our heart condemn us not, then have we confidence toward God (1John 3:21)` + bcolors.ENDC)
    case 947:
        bcolors.Colors(`(📜Greater is he that is in you, than he that is in the world (1John 4:4)` + bcolors.ENDC)
    case 948:
        bcolors.Colors(`(📜Let us love one another: for love is of God (1John 4:7)` + bcolors.ENDC)
    case 949:
        bcolors.Colors(`(📜God is love (1John 4:8)` + bcolors.ENDC)
    case 950:
        bcolors.Colors(`(📜The love of God was manifested toward us (1John 4:9)` + bcolors.ENDC)
    case 951:
        bcolors.Colors(`(📜He loved us, & sent his Son to be the propitiation for our sins (1John 4:10)` + bcolors.ENDC)
    case 952:
        bcolors.Colors(`(📜if God so loved us, we ought also to love one another (1John 4:11)` + bcolors.ENDC)
    case 953:
        bcolors.Colors(`(📜The Father sent the Son to be the Saviour of the world (1John 4:14)` + bcolors.ENDC)
    case 954:
        bcolors.Colors(`(📜& we have known & believed the love that God hath to us (1John 4:16)` + bcolors.ENDC)
    case 955:
        bcolors.Colors(`(📜There is no fear in love; but perfect love casteth out fear (1John 4:18)` + bcolors.ENDC)
    case 956:
        bcolors.Colors(`(📜We love Him, because he first loved us (1John 4:19)` + bcolors.ENDC)
    case 957:
        bcolors.Colors(`(📜His commandments are not grievous (1John 5:3)` + bcolors.ENDC)
    case 958:
        bcolors.Colors(`(📜Whatsoever is born of God overcometh the world (1John 5:4)` + bcolors.ENDC)
    case 959:
        bcolors.Colors(`(📜This is the victory that overcometh the world, even our faith. (1John 5:4)` + bcolors.ENDC)
    case 960:
        bcolors.Colors(`(📜God hath given to us eternal life (1John 5:11)` + bcolors.ENDC)
    case 961:
        bcolors.Colors(`(📜That ye may know that ye have eternal life (1John 5:13)` + bcolors.ENDC)
    case 962:
        bcolors.Colors(`(📜He that is begotten of God keepeth himself, & that wicked one toucheth him not (1John 5:18)` + bcolors.ENDC)
    case 963:
        bcolors.Colors(`(📜We know that we are of God, & the whole world lieth in wickedness (1John 5:19)` + bcolors.ENDC)
    case 964:
        bcolors.Colors(`(📜The Son of God is come, & hath given us an underst&ing (1John 5:20)` + bcolors.ENDC)
    case 965:
        bcolors.Colors(`(📜This is love, that we walk after His commandments (2John 1:6)` + bcolors.ENDC)
    case 966:
        bcolors.Colors(`(📜I wish above all things that thou mayest prosper & be in health (3John 1:2)` + bcolors.ENDC)
    case 967:
        bcolors.Colors(`(📜Keep yourselves in the love of God (Jude 1:21)` + bcolors.ENDC)
    case 968:
        bcolors.Colors(`(📜Blessed is he that readeth, & they that hear the words of this prophecy (Rev.1:3)` + bcolors.ENDC)
    case 969:
        bcolors.Colors(`(📜Behold, He cometh with clouds; & every eye shall see Him (Rev.1:7)` + bcolors.ENDC)
    case 970:
        bcolors.Colors(`(📜I am Alpha & Omega, the beginning & the ending, saith the Lord (Rev.1:8)` + bcolors.ENDC)
    case 971:
        bcolors.Colors(`(📜I know thy works, & thy labour, & thy patience (Rev.2:2)` + bcolors.ENDC)
    case 972:
        bcolors.Colors(`(📜To him that overcometh will I give to eat of the tree of life (Rev.2:7)` + bcolors.ENDC)
    case 973:
        bcolors.Colors(`(📜Be thou faithful unto death, & I will give thee a crown of life (Rev.2:10)` + bcolors.ENDC)
    case 974:
        bcolors.Colors(`(📜To him that overcometh will I give to eat of the hidden manna (Rev.2:17)` + bcolors.ENDC)
    case 975:
        bcolors.Colors(`(📜Remember therefore how thou hast received & heard, & hold fast, & repent (Rev.3:3)` + bcolors.ENDC)
    case 976:
        bcolors.Colors(`(📜Behold, I have set before thee an open door, & no man can shut it (Rev.3:8)` + bcolors.ENDC)
    case 978:
        bcolors.Colors(`(📜Behold, I come quickly! (Rev.3:11)` + bcolors.ENDC)
    case 979:
        bcolors.Colors(`(📜Hold that fast which thou hast, that no man take thy crown (Rev.3:11)` + bcolors.ENDC)
    case 980:
        bcolors.Colors(`(📜Him that overcometh will I make a pillar in the temple of My God (Rev.3:12)` + bcolors.ENDC)
    case 981:
        bcolors.Colors(`(📜He that hath an ear, let him hear what the Spirit saith unto the churches (Rev.3:13)` + bcolors.ENDC)
    case 982:
        bcolors.Colors(`(📜As many as I love, I rebuke & chasten (Rev.3:19)` + bcolors.ENDC)
    case 983:
        bcolors.Colors(`(📜Behold, I st& at the door, & knock (Rev.3:20)` + bcolors.ENDC)
    case 984:
        bcolors.Colors(`(📜If any man hear my voice, & open the door, I will come in to him (Rev.3:20)` + bcolors.ENDC)
    case 985:
        bcolors.Colors(`(📜To him that overcometh will I grant to sit with Me in My throne (Rev.3:21)` + bcolors.ENDC)
    case 986:
        bcolors.Colors(`(📜Holy, holy, holy, LORD God Almighty (Rev.4:8)` + bcolors.ENDC)
    case 987:
        bcolors.Colors(`(📜Thou art worthy, O Lord, to receive glory & honour & power (Rev.4:11)` + bcolors.ENDC)
    case 988:
        bcolors.Colors(`(📜& hast made us unto our God kings & priests (Rev.5:10)` + bcolors.ENDC)
    case 989:
        bcolors.Colors(`(📜Worthy is the Lamb that was slain (Rev.5:12)` + bcolors.ENDC)
    case 990:
        bcolors.Colors(`(📜Salvation to our God which sitteth upon the throne, & unto the Lamb! (Rev.7:10)` + bcolors.ENDC)
    case 991:
        bcolors.Colors(`(📜& God shall wipe away all tears from their eyes (Rev.7:17)` + bcolors.ENDC)
    case 992:
        bcolors.Colors(`(📜& the smoke of the incense, which came with the prayers of the saints, ascended up before God (Rev.8:4)` + bcolors.ENDC)
    case 993:
        bcolors.Colors(`(📜They overcame him by the blood of the Lamb, & by the word of their testimony (Rev.12:11)` + bcolors.ENDC)
    case 994:
        bcolors.Colors(`(📜Worship Him that made heaven, & earth, & the sea, & the fountains of waters (Rev.14:7)` + bcolors.ENDC)
    case 995:
        bcolors.Colors(`(📜Great & marvellous are Thy works, Lord God Almighty! (Rev.15:3)` + bcolors.ENDC)
    case 996:
        bcolors.Colors(`(📜Who shall not fear Thee, O Lord, & glorify Thy name? (Rev.15:4)` + bcolors.ENDC)
    case 997:
        bcolors.Colors(`(📜Thou art righteous, O Lord, which art, & wast, & shalt be (Rev.16:5)` + bcolors.ENDC)
    case 998:
        bcolors.Colors(`(📜Even so, Lord God Almighty, true & righteous are thy judgments (Rev.16:7)` + bcolors.ENDC)
    case 999:
        bcolors.Colors(`(📜Blessed is he that watcheth, & keepeth his garments (Rev.16:15)` + bcolors.ENDC)
    case 1000:
        bcolors.Colors(`(📜Behold, I make all things new (Rev.21:5)` + bcolors.ENDC)
    case 1001:
        bcolors.Colors(`(📜I will give unto him that is athirst of the fountain of the water of life freely (Rev.21:6)` + bcolors.ENDC)
    case 1002:
        bcolors.Colors(`(📜He that overcometh shall inherit all things (Rev.21:7)` + bcolors.ENDC)
    case 1003:
        bcolors.Colors(`(📜Blessed is he that keepeth the sayings of the prophecy of this book (Rev.22:7)` + bcolors.ENDC)
    case 1004:
        bcolors.Colors(`(📜He that is holy, let him be holy still (Rev.22:11)` + bcolors.ENDC)
    case 1005:
        bcolors.Colors(`(📜Whosoever will, let him take the water of life freely (Rev.22:17)` + bcolors.ENDC)
    case 1006:
        bcolors.Colors(`(📜Even so, come, Lord Jesus! (Rev.22:20)` + bcolors.ENDC)
    }
}

func Verse() {
    Salvation := &Scriptures{}
    Salvation.TheWord()
}

