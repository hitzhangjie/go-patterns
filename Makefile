all:
	#go work init
	echo "update go.work"
	for i in `find -iname go.mod`; do echo ${i%%/go.mod}; go work use "${i%%/go.mod}"; done

.PHONY: all
