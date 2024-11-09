## Validators is the lowest level package for STY-Holdings.

## Restrictions:
Validators can not call out to any other package because it may create a circular reference.
This means that code can be duplicated in the package that also resides in another package.  