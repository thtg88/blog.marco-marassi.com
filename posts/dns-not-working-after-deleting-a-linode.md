---
title: 'DNS Not Working After Deleting A Linode'
date: '2020-05-07'
---

Over the past years, I always managed my web projects that required a server-side architecture with [Linode](https://www.linode.com/), which provides a great service for a very small cost for my hobby web dev projects.

Together with providing bare-metal servers, they also provide a domain management feature, which I've been using to keep things centralised and easy to manage, allowing me not to bounce between dozen of sites to just put a simple project online.

In the last few days, I've found myself moving some stuff around hosting accounts between Linode and [Heroku](https://www.heroku.com/).

While doing so, as my running linode server instance had no applications running on it, I've decided to shut it down and delete it, as it seemed like a waste of money leaving the server running for nothing.

After finishing moving my projects and having pointed the DNS back at Heroku, it was taking suspiciously long for the DNS to propagate. The infamous `DNS_PROBE_FINISHED_NXDOMAIN` started to creep up at me on my browser which raised my suspicion, and I started `dig`-ing around the reason of this unusual delay.

A quick `dig marco-marassi.com` resulted in:

```
; <<>> DiG 9.10.6 <<>> marco-marassi.com
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NXDOMAIN, id: 6000
;; flags: qr rd ra; QUERY: 1, ANSWER: 0, AUTHORITY: 1, ADDITIONAL: 1

;; OPT PSEUDOSECTION:
; EDNS: version: 0, flags:; udp: 512
;; QUESTION SECTION:
;marco-marassi.com. IN A

;; AUTHORITY SECTION:
com. 621 IN SOA a.gtld-servers.net. nstld.verisign-grs.com. 1588937086 1800 900 604800 86400

;; Query time: 395 msec
;; SERVER: 192.168.1.1#53(192.168.1.1)
;; WHEN: Fri May 08 12:29:56 BST 2020
;; MSG SIZE rcvd: 119
```

Which means the DNS authoritative server of the `.com` domain (`gtld-servers.net`) was replying to my query, instead of my Linode nameservers, or my domain registrars.

After trying `nslookup` with no avail (`server can't find marco-marassi.com: NXDOMAIN`), I thought maybe it was something to do with my network or router doing some sort of weird caching, so I reverted to my phone, and the good old [MxToolbox SuperTool](https://mxtoolbox.com/SuperTool.aspx).

Even MxToolbox was telling me that there were no records for it.
I started trying some sub-domains of other projects I had running on the `marco-marassi.com` domain, but again, no luck.
The domain seemed to have vanished!

After a brief moment of panic, I travelled back in my mind to what I had done in the past half hour trying to think what could have caused this sudden DNS outage. I had not touch my domain registrar panel nor my nameservers, so there was no reason for it to just disappear from the internet.

I had deleted my linode server! But that surely could not be it, it had nothing to do with the DNS so why would that affect it?

After trying pointing my DNS pretty much anywhere to get a positive outcome, I decided to turn to Linode's great support service.

After a some time they did confirm that:
> In order to continue serving this domain, you would need to add a Linode to your account or move your DNS records to your registrar or another location. There's a blurb about this requirement in our DNS Manager guide.

Even with the [DNS Manager guide](https://www.linode.com/docs/platform/manager/dns-manager/) at hand, I was struggling to find it, but finally out of the blue, the infamous note:
> To use the Linode DNS Manager to serve your domains, you must have an active Linode on your account. If you remove all active Linodes, your domains will no longer be served.

That was it! I quickly spinned up a new linode Nanode instance, and voilà! All my projects were back online! 🤦‍
