# ban_ip


```
firewall-cmd --permanent --zone=public --new-ipset=blacklist --type=hash:ip
firewall-cmd --permanent --zone=public --add-rich-rule='rule source ipset=blacklist drop'
firewall-cmd --permanent --zone=public --ipset=blacklist --add-entry=222.222.222.222
firewall-cmd --reload
```

