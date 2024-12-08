
# General Purpose Time Entries Text Format

```
YYYY-MM-DD
  HH:MM comment #tag1 #tag2
comment
  comment
  HH:MM comment #tag1 #tag3

YYYY-MM-DD
  HH:MM comment #tag1 #break
```

Reserved keyword tag `#break`


- Build a parser? More stable parsing of files, flexibility in format
- Use templates for output? (e.g. `stpw list --format=json | jq [].duration`)
