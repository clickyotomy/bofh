package main

// Excuses is a list of excuses from BOFH.
// Source: http://pages.cs.wisc.edu/~ballard/bofh/excuses.
var Excuses []string = []string{
	"Clock speed.",
	"Solar Flares.",
	"Electromagnetic radiation from satellite debris.",
	"Static from nylon underwear.",
	"Static from plastic slide rules.",
	"Global warming.",
	"Poor power conditioning.",
	"Static build-up.",
	"Doppler effect.",
	"Hardware stress fractures.",
	"Magnetic interference from money/credit cards.",
	"Dry joints on cable plug.",
	"We're waiting for [the phone company] to fix that line.",
	"Sounds like a Windows problem, try calling Microsoft support.",
	"Temporary routing anomaly.",
	"Somebody was calculating 'pi' on the server.",
	"Fat electrons in the lines.",
	"Excess surge protection.",
	"Floating point processor overflow.",
	"Divide-by-zero error.",
	"POSIX compliance problem.",
	"Monitor resolution too high.",
	"Improperly oriented keyboard.",
	"Network packets travelling up-hill (use a carrier pigeon).",
	"Decreasing electron flux.",
	"First Saturday after first full moon in winter.",
	"Radiosity depletion.",
	"CPU radiator broken.",
	"It works the way the Wang did, what's the problem.",
	"Positron router malfunction.",
	"Cellular telephone interference.",
	"Techtonic stress.",
	"Piezoelectric interference.",
	"(L)user error.",
	"Working as designed.",
	"Dynamic software linking table corrupted.",
	"Heavy gravity fluctuation, move computer to floor rapidly!",
	"Secretary plugged hairdryer into UPS.",
	"Terrorist activities.",
	"Not enough memory, go get system upgrade.",
	"Interrupt configuration error.",
	"Spaghetti cable cause packet failure.",
	"Boss forgot system password.",
	"Bank Holiday: System operating credits not recharged.",
	"Virus attack, (l)user responsible.",
	"Waste water tank overflowed onto computer.",
	"Complete Transient Lockout.",
	"Bad ether in the cables.",
	"Bogon emissions.",
	"Change in Earth's rotational speed.",
	"Cosmic ray particles crashed through the hard disk platter.",
	"Smell from unhygienic janitorial staff wrecked the tape heads.",
	"Little hamster in running wheel had coronary; waiting for replacement " +
		"to be Fedexed from Wyoming.",
	"Evil dogs hypnotised the night shift.",
	"Plumber mistook routing panel for decorative wall fixture.",
	"Electricians made popcorn in the power supply.",
	"Groundskeepers stole the root password.",
	"High pressure system failure.",
	"Failed trials, system needs redesign.",
	"System has been recalled.",
	"Not approved by the FCC.",
	"Need to wrap system in aluminum foil to fix problem.",
	"Not properly grounded, please bury computer.",
	"CPU needs recalibration.",
	"System needs to be rebooted.",
	"Bit bucket overflow.",
	"Descramble code needed from software company.",
	"Only available on a need to know basis.",
	"Knot in cables caused data stream to become twisted and kinked.",
	"Nesting roaches shorted out the ether cable.",
	"The file system is full of it.",
	"Satan did it.",
	"Daemons did it.",
	"You're out of memory.",
	"There isn't any problem.",
	"Unoptimized hard drive.",
	"Typo in the code.",
	"Yes, yes, its called a design limitation.",
	"Look, buddy: Windows 3.1 _is_ a General Protection Fault.",
	"That's a great computer you have there; have you considered how it " +
		"would work as a BSD machine?",
	"Please excuse me, I have to circuit an AC line through my head to get " +
		"this database working.",
	"Yeah, yo mama dresses you funny and you need a mouse to delete files.",
	"Support staff hung over, send aspirin and come back LATER.",
	"Someone is standing on the ethernet cable, causing a kink in the cable.",
	"Windows 95 undocumented 'feature'.",
	"Runt packets.",
	"Password is too complex to decrypt.",
	"Boss' kid fucked up the machine.",
	"Electromagnetic energy loss.",
	"Budget cuts.",
	"Mouse chewed through power cable.",
	"Stale file handle (next time use Tupperware (TM)!).",
	"Feature not yet implemented.",
	"Internet outage.",
	"Pentium FDIV bug.",
	"Vendor no longer supports the product.",
	"Small animal kamikaze attack on power supplies.",
	"The vendor put the bug there.",
	"SIMM crosstalk.",
	"IRQ dropout.",
	"Collapsed Backbone.",
	"Power company testing new voltage spike (creation) equipment.",
	"Operators on strike due to broken coffee machine.",
	"Backup tape overwritten with copy of system manager's favourite CD.",
	"UPS interrupted the server's power.",
	"The electrician didn't know what the yellow cable was so he yanked " +
		"the ethernet out.",
	"The keyboard isn't plugged in.",
	"The air conditioning water supply pipe ruptured over the machine room.",
	"The electricity substation in the car park blew up.",
	"The rolling stones concert down the road caused a brown out.",
	"The salesman drove over the CPU board.",
	"The monitor is plugged into the serial port.",
	"Root nameservers are out of sync.",
	"Electro-magnetic pulses from French above ground nuke testing.",
	"Your keyboard's space bar is generating spurious keycodes.",
	"The real TTYs became pseudo TTYs and vice-versa.",
	"The printer thinks its a router.",
	"The router thinks its a printer.",
	"Evil hackers from Serbia.",
	"We just switched to FDDI.",
	"Halon system went off and killed the operators.",
	"Because Bill Gates is a Jehovah's witness and so nothing can work " +
		"on St. Swithin's day.",
	"User to computer ratio too high.",
	"User to computer ration too low.",
	"We just switched to Sprint.",
	"It has Intel Inside.",
	"Sticky bits on disk.",
	"Power company having EMP problems with their reactor.",
	"The ring needs another token.",
	"New management.",
	"telnet: Unable to connect to remote host: Connection refused.",
	"SCSI Chain overterminated.",
	"It's not plugged in.",
	"Because of network lag due to too many people playing deathmatch.",
	"You put the disk in upside down.",
	"Daemons loose in system.",
	"User was distributing pornography on server; system seized by FBI.",
	"BNC (Brain Not Connected)",
	"UBNC (User Brain Not Connected)",
	"LBNC ((L)user Brain Not Connected)",
	"Disks spinning backwards - toggle the hemisphere jumper.",
	"New guy cross-connected phone lines with AC power bus.",
	"Had to use hammer to free stuck disk drive heads.",
	"Too few computrons available.",
	"Flat tire on station wagon with tapes. (\"Never underestimate the " +
		"bandwidth of a station wagon full of tapes hurling down the " +
		"highway\" - Andrew S. Tannenbaum)",
	"Communications satellite used by the military for Star Wars.",
	"Party-bug in the Aloha protocol.",
	"Insert coin for new game.",
	"Dew on the telephone lines.",
	"Arcserve crashed the server again.",
	"Some one needed the powerstrip, so they pulled the switch plug.",
	"My pony-tail hit the on/off switch on the power strip.",
	"Big to little endian conversion error.",
	"You can tune a file system, but you can't tune a fish " +
		"(from most `tunefs' man pages).",
	"Dumb terminal.",
	"Zombie processes haunting the computer.",
	"Incorrect time synchronization.",
	"Defunct processes.",
	"Stubborn processes.",
	"Non-redundant fan failure.",
	"Monitor VLF leakage",
	"Bugs in the RAID.",
	"No 'any' key on keyboard.",
	"Root Rot.",
	"Backbone Scoliosis.",
	"/pub/lunch",
	"Excessive collisions and not enough packet ambulances.",
	"le0: no carrier: transceiver cable problem?",
	"Broadcast packets on wrong frequency.",
	"Popper unable to process jumbo kernel",
	"NOTICE: alloc: /dev/null: filesystem full",
	"Pseudo-user on a pseudo-terminal.",
	"Recursive traversal of loopback mount points.",
	"Backbone adjustment.",
	"OS swapped to disk.",
	"Vapors from evaporating sticky-note adhesives.",
	"Sticktion.",
	"Short leg on process table.",
	"Multicasts on broken packets.",
	"Ether leak!",
	"Atilla the Hub.",
	"Endothermal recalibration.",
	"Filesystem not big enough for Jumbo Kernel Patch.",
	"Loop found in loop in redundant loopback.",
	"System consumed all the paper for paging.",
	"Permission denied.",
	"Reformatting Page. Wait...",
	"...disk or the processor is on fire.",
	"SCSI's too wide.",
	"Proprietary Information.",
	"Just type 'mv * /dev/null'.",
	"Runaway cat on system.",
	"Did you pay the new Support Fee?",
	"We only support a 1200 bps connection.",
	"We only support a 28000 bps connection.",
	"Me no Internet, only janitor, me just wax floors.",
	"I'm sorry a Pentium won't do, you need an SGI to connect with us.",
	"Post-it Note sludge leaked into the monitor.",
	"The curls in your keyboard cord are losing electricity.",
	"The monitor needs another box of pixels.",
	"RPC_PMAP_FAILURE",
	"kernel panic: write-only-memory (/dev/wom0) capacity exceeded.",
	"Write-only-memory subsystem too slow for this machine. " +
		"Contact your local dealer.",
	"Just pick up the phone and give modem connect sounds. " +
		"\"Well you said we should get more lines so we don't " +
		"have voice lines.\"",
	"Quantum dynamics are affecting the transistors.",
	"Police are examining all the Internet packets in the search " +
		"for a narco-net-trafficker.",
	"We are currently trying a new concept of using a live mouse. " +
		"Unfortunately, one has yet to survive being hooked up to the " +
		"computer... Please bear with us.",
	"Your mail is being routed through Germany... and they're censoring us.",
	"Only people with names beginning with 'A' are getting mail this " +
		"week (a la Microsoft).",
	"We didn't pay the Internet bill and it's been cut off.",
	"Lightning strikes.",
	"Of course it doesn't work. We've performed a software upgrade.",
	"Change your language to Finnish.",
	"Fluorescent lights are generating negative ions. If turning them off " +
		"doesn't work, take them out and put tin foil on the ends.",
	"High nuclear activity in your area.",
	"What office are you in? Oh, that one. Did you know that your building " +
		"was built over the universities first nuclear research site? " +
		"And wow, aren't you the lucky one? Your office is right over " +
		"where the core is buried!",
	"The MGs ran out of gas.",
	"The UPS doesn't have a battery backup.",
	"Recursivity. Call back if it happens again.",
	"Someone thought The Big Red Button was a light switch.",
	"The mainframe needs to rest. It's getting old, you know.",
	"I'm not sure. Try calling the Internet's head " +
		"office -- it's in the book.",
	"The lines are all busy (busied out, that is -- why let " +
		"them in to begin with?).",
	"Jan 9 16:41:27 huber su: 'su root' succeeded for [...] on /dev/pts/1",
	"It's those computer people in X {city of world}. " +
		"They keep stuffing things up.",
	"A Star Wars satellite accidently blew up the WAN.",
	"Fatal error right in front of screen.",
	"That function is not currently supported, but Bill Gates assures " +
		"us it will be featured in the next upgrade.",
	"Wrong polarity of neutron flow.",
	"(L)users learning curve appears to be fractal.",
	"We had to turn off that service to comply with the CDA Bill.",
	"Ionization from the air-conditioning.",
	"TCP/IP UDP alarm threshold is set too low.",
	"Someone is broadcasting pygmy packets and the router doesn't " +
		"know how to deal with them.",
	"The new frame relay network hasn't bedded down the software " +
		"loop transmitter yet.",
	"Fanout dropping voltage too much, try cutting some of those " +
		"little traces.",
	"Plate voltage too low on demodulator tube.",
	"You did wha... Oh _dear_ ...",
	"CPU needs bearings repacked.",
	"Too many little pins on CPU confusing it, bend back and forth until " +
		"10-20% are neatly removed. Do _not_ leave metal bits visible!",
	"_Rosin_ core solder? But...",
	"Software uses U.S. measurements, but the OS is in metric...",
	"The computer fleetly, mouse and all.",
	"Your cat tried to eat the mouse.",
	"The Borg tried to assimilate your system. Resistance is futile.",
	"It must have been the lightning storm we had " +
		"yesterday/last week/last month.",
	"Due to Federal Budget problems we have been forced to cut back " +
		"on the number of users able to access the system at one time. " +
		"(namely none allowed...)",
	"Too much radiation coming from the soil.",
	"Unfortunately we have run out of bits/bytes/whatever. " +
		"Don't worry, the next supply will be coming next week.",
	"Program load too heavy for processor to lift.",
	"Processes running slowly due to weak power supply.",
	"Our ISP is having {switching,routing,SMDS,frame relay} problems.",
	"We've run out of licenses.",
	"Interference from lunar radiation.",
	"Standing room only on the bus.",
	"You need to install an RTFM interface.",
	"That would be because the software doesn't work.",
	"That's easy to fix, but I can't be bothered.",
	"Someone's tie is caught in the printer, and if anything " +
		"else gets printed, he'll be in it too.",
	"We're upgrading /dev/null.",
	"The Usenet news is out of date.",
	"Our POP server was kidnapped by a weasel.",
	"It's stuck in the Web.",
	"Your modem doesn't speak English.",
	"The mouse escaped.",
	"All of the packets are empty.",
	"The UPS is on strike.",
	"Neutrino overload on the nameserver.",
	"Melting hard drives.",
	"Someone has messed up the kernel pointers.",
	"The kernel license has expired.",
	"Netscape has crashed.",
	"The cord jumped over and hit the power switch.",
	"It was OK before you touched it.",
	"Bit Rot.",
	"U.S. Postal Service.",
	"Your Flux Capacitor has gone bad.",
	"The Dilithium Crystals need to be rotated.",
	"The static electricity routing is acting up...",
	"Traceroute says that there is a routing problem in the backbone. " +
		"It's not our problem.",
	"The co-locator cannot verify the frame-relay gateway to the ISDN server.",
	"High altitude condensation from U.S.A.F. prototype aircraft " +
		"has contaminated the primary subnet mask. Turn off your computer " +
		"for 9 days to avoid damaging it.",
	"Lawn mower blade in your fan need sharpening.",
	"Electrons on a bender.",
	"Telecommunications is upgrading.",
	"Telecommunications is downgrading.",
	"Telecommunications is downshifting.",
	"Hard drive sleeping. Let it wake up on it's own...",
	"Interference between the keyboard and the chair.",
	"The CPU has shifted, and become decentralized.",
	"Due to the CDA, we no longer have a root account.",
	"We ran out of dial tone and we're and waiting for the phone company " +
		"to deliver another bottle.",
	"You must've hit the wrong 'any' key.",
	"PCMCIA slave driver.",
	"The token fell out of the ring. Call us when you find it.",
	"The hardware bus needs a new token.",
	"Too many interrupts.",
	"Not enough interrupts.",
	"The data on your hard drive is out of balance.",
	"Digital Manipulator exceeding velocity parameters",
	"Appears to be a slow/narrow SCSI-0 Interface problem",
	"Micro-electronic Riemannian curved-space fault in " +
		"write-only file system.",
	"Fractal radiation jamming the backbone.",
	"Routing problems on the neural net.",
	"IRQ-problems with the Un-Interruptible-Power-Supply.",
	"CPU-angle has to be adjusted because of vibrations coming " +
		"from the nearby road.",
	"Emissions from GSM-phones",
	"CD-ROM server needs recalibration.",
	"Firewall needs cooling.",
	"Asynchronous inode failure.",
	"Transient bus protocol violation.",
	"Incompatible bit-registration operators.",
	"Your process is not ISO 9000 compliant.",
	"You need to upgrade your VESA local bus to a MasterCard local bus.",
	"The recent proliferation of nuclear testing.",
	"Elves on strike. (Why do they call EMAG Elf Magic?)",
	"Internet exceeded (L)user level, please wait until a (l)user logs " +
		"off before attempting to log back on.",
	"Your e-mail is now being delivered by the USPS.",
	"Your computer hasn't been returning all the bits it gets " +
		"from the Internet.",
	"You've been infected by the Telescoping Hubble virus.",
	"Scheduled global CPU outage.",
	"Your Pentium has a heating problem - try cooling it with " +
		"ice cold water. (Do not turn off your computer, you do not " +
		"want to cool down the Pentium Chip while he isn't working, do you?)",
	"Your processor has processed too many instructions. " +
		"Turn it off immediately, do not type any commands!",
	"Your packets were eaten by the Terminator.",
	"Your processor does not develop enough heat.",
	"We need a licensed electrician to replace the light " +
		"bulbs in the computer room.",
	"The POP server is out of Coke.",
	"Fiber optics caused gas main leak.",
	"Server depressed, needs Prozac.",
	"Quantum Decoherence.",
	"Those damn raccoons!",
	"Suboptimal routing experience.",
	"A plumber is needed, the network drain is clogged.",
	"50% of the manual is in '.pdf' README files.",
	"The AA battery in the wallclock sends magnetic interference.",
	"The XY axis in the trackball is coordinated with the summer solstice.",
	"The butane lighter causes the pin-cushioning.",
	"Old inkjet cartridges emanate barium-based fumes.",
	"Manager in the cable duct.",
	"We'll fix that in the next (upgrade/update/patch release/service pack).",
	"HTTPD Error 666: BOFH was here.",
	"HTTPD Error 4004: Very old Intel CPU - insufficient processing power.",
	"The ATM board has run out of 10 pound notes. We are having a " +
		"whip round to refill it, care to contribute?",
	"Network Failure - call NBC.",
	"Having to manually track the satellite.",
	"Your/our computer(s) had suffered a memory leak, and we are " +
		"waiting for them to be topped up.",
	"The rubber band broke.",
	"We're on Token Ring, and it looks like the token got loose.",
	"Stray alpha particles from memory packaging caused hard " +
		"memory error on server.",
	"Paradigm shift... without a clutch.",
	"PEBKAC (Problem Exists Between Keyboard And Chair)",
	"The cables are not the same length.",
	"Second-system effect.",
	"Chewing gum on /dev/sd3c.",
	"Boredom in the kernel.",
	"The daemons! The daemons! The terrible daemons!",
	"I'd love to help you -- it's just that the boss won't let " +
		"me near the computer.",
	"Struck by the Good Times virus",
	"YOU HAVE AN I/O ERROR: Incompetent Operator error.",
	"Your parity check is overdrawn and you're out of cache.",
	"Communist revolutionaries taking over the server room and demanding " +
		"all the computers in the building or they shoot the sys-admin. " +
		"Poor misguided fools.",
	"Plasma conduit breach.",
	"Out of cards on drive 'D:'.",
	"Sand fleas eating the Internet cables.",
	"Parallel processors running perpendicular today.",
	"ATM cell has no roaming feature turned on, notebooks can't connect.",
	"Webmasters kidnapped by evil cult.",
	"Failure to adjust for daylight savings time.",
	"Virus transmitted from computer to sysadmins.",
	"Virus due to computers having unsafe sex.",
	"Incorrectly configured static routes on the core-routers.",
	"Forced to support NT servers; sys-admins quit.",
	"Suspicious pointer corrupted virtual machine.",
	"It's the InterNIC's fault.",
	"Root name servers corrupted.",
	"Budget cuts forced us to sell all the power cords for the servers.",
	"Someone hooked the twisted pair wires into the answering machine.",
	"Operators killed by year 2000 bug bite.",
	"We've picked COBOL as the language of choice.",
	"Operators killed when huge stack of backup tapes fell over.",
	"Robotic tape changer mistook operator's tie for a backup tape.",
	"Someone was smoking in the computer room and set off the halon systems.",
	"Your processor has taken a ride to Heaven's Gate on the UFO " +
		"behind Hale-Bopp's comet.",
	"It's an ID-10-T error.",
	"Dyslexics retyping hosts file on servers.",
	"The Internet is being scanned for viruses.",
	"Your computer's union contract is set to expire at midnight.",
	"Bad user karma.",
	"/dev/clue was linked to /dev/null",
	"Increased sunspot activity.",
	"We already sent around a notice about that.",
	"It's union rules. There's nothing we can do about it. Sorry.",
	"Interference from the Van Allen Belt.",
	"Jupiter is aligned with Mars.",
	"Redundant ACLs.",
	"Mail server hit by UniSpammer.",
	"T-1's congested due to porn traffic to the news server.",
	"Data for intranet got routed through the extranet and " +
		"landed on the Internet.",
	"We are a 100% Microsoft Shop.",
	"We are Microsoft. What you are experiencing is not a problem; " +
		"it is an undocumented feature.",
	"Sales staff sold a product we don't offer.",
	"Secretary sent chain letter to all 5000 employees.",
	"Sysadmin didn't hear pager go off due to loud music from " +
		"bar-room speakers.",
	"Sysadmin accidentally destroyed pager with a large hammer.",
	"Sysadmins unavailable because they are in a meeting talking " +
		"about why they are unavailable so much.",
	"Bad cafeteria food landed all the sysadmins in the hospital.",
	"Route flapping at the NAP.",
	"Computers under water due to SYN flooding.",
	"The vulcan-death-grip ping has been applied.",
	"Electrical conduits in machine room are melting.",
	"Traffic jam on the Information Superhighway.",
	"Radial Telemetry Infiltration.",
	"Cow-tippers tipped a cow onto the server.",
	"Tachyon emissions overloading the system.",
	"Maintenance window broken.",
	"We're out of slots on the server.",
	"Computer room being moved. Our systems are down for the weekend.",
	"Sysadmins busy fighting spam.",
	"Repeated reboots of the system failed to solve problem.",
	"Feature was not beta tested.",
	"Domain controller not responding.",
	"Someone else stole your IP address, call the Internet detectives!",
	"It's not RFC-822 compliant.",
	"Operation failed because there is no message for this error (#1014).",
	"Stop bit received",
	"The Internet is needed to catch the etherbunny.",
	"Network down, IP packets delivered via UPS.",
	"Firmware update in the coffee machine.",
	"Temporal anomaly.",
	"Mouse has out-of-cheese-error.",
	"Borg implants are failing.",
	"Borg nanites have infested the server.",
	"error: one bad user found in front of screen",
	"Please state the nature of the technical emergency",
	"Internet shut down due to maintenance.",
	"Daemon escaped from pentagram.",
	"Crop circles in the corn shell.",
	"Sticky bit has come loose.",
	"Hot Java has gone cold.",
	"Cache miss - please take better aim next time.",
	"Hash table has woodworm.",
	"Trojan horse ran out of hay.",
	"Zombie processes detected, machine is haunted.",
	"Overflow error in /dev/null.",
	"Browser's cookie is corrupted -- someone's been nibbling on it.",
	"Mailer-daemon is busy burning your message in hell.",
	"According to Microsoft, it's by design.",
	"'vi' needs to be upgraded to 'vii'.",
	"Greenpeace free()'d the mallocs",
	"Terrorists crashed an airplane into the server room, " +
		"have to remove /bin/laden. (rm -rf /bin/laden).",
	"Astropneumatic oscillations in the water-cooling.",
	"Somebody ran the operating system through a spelling checker.",
	"Rhythmic variations in the voltage reaching the power supply.",
	"Keyboard Actuator Failure. Order and Replace.",
	"Packet held up at customs.",
	"Propagation delay.",
	"High line impedance.",
	"Someone set us up the bomb.",
	"Power surges on the Underground.",
	"Don't worry; it's been deprecated. The new one is worse.",
	"Excess condensation in cloud network.",
	"It is a layer 8 problem.",
	"The math co-processor had an overflow error that leaked out and " +
		"shorted the RAM.",
	"Leap second overloaded RHEL6 servers.",
	"DNS server drank too much and had a hiccup.",
	"Your machine had the fuses in backwards.",
}
