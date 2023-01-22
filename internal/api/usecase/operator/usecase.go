//go:generate mockgen -destination=mock/usecase_$GOPACKAGE.gen.go -package=usecase_$GOPACKAGE github.com/basslove/daradara/internal/api/usecase/$GOPACKAGE GetSightCategoriesInputPort,GetSightGenresInputPort,PostOperatorsInputPort,PostOperatorsSignInInputPort

package operator
