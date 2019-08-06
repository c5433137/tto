package go_

var (
	// <R> : 仓储类的名称
	// <R2> : 函数后添加的仓储类名称
	// <E> : 实体
	// <E2> : 包含包名的实体
	// t : 仓库类对象引用
	TPL_ENTITY_REP = `
            package repo

/**
 * this file is auto generated by tto v{{.global.Version}} !
 * if you want to modify this code,please read guide doc
 * and modify code template later.
 *
 * guide please see https://github.com/ixre/tto
 * generate time: {{.global.Time}}*
 */
#target!{{.global.Pkg}}/repo/{{.table.Name}}_repo.go

            import(
                "log"
                "{{.global.Pkg}}/model"
                "{{.global.Pkg}}/ifce"
                "database/sql"
                "github.com/ixre/gof/db/orm"
            )

            // Create new {{.table.Title}}Repo
            func New{{.table.Title}}Repo(o orm.Orm)*{{.table.Title}}Repo{
                return &{{.table.Title}}Repo{
                    _orm:o,
                }
            }
			var _ ifce.I{{.table.Title}}Repo = new({{.table.Title}}RepoImpl)
            type {{.table.Title}}RepoImpl struct{
                _orm orm.Orm
            }

            // Get {{.table.Comment}}
            func (t *{{.table.Title}}RepoImpl) Get(primary interface{})*model.{{.table.Title}}{
                e := model.{{.table.Title}}{}
                err := t._orm.Get(primary,&e)
                if err == nil{
                    return &e
                }
                if err != sql.ErrNoRows{
                  log.Println("[ Orm][ Error]:",err.Error(),"; Entity:{{.table.Title}}")
                }
                return nil
            }

            // GetBy {{.table.Comment}}
            func (t *{{.table.Title}}RepoImpl) GetBy(where string,v ...interface{})*model.{{.table.Title}}{
                e := model.{{.table.Title}}{}
                err := t._orm.GetBy(&e,where,v...)
                if err == nil{
                    return &e
                }
                if err != sql.ErrNoRows{
                  log.Println("[ Orm][ Error]:",err.Error(),"; Entity:{{.table.Title}}")
                }
                return nil
            }

            // Select {{.table.Comment}}
            func (t *{{.table.Title}}RepoImpl) Select(where string,v ...interface{})[]*model.{{.table.Title}} {
                list := make([]*model.{{.table.Title}},0)
                err := t._orm.Select(&list,where,v...)
                if err != nil && err != sql.ErrNoRows{
                  log.Println("[ Orm][ Error]:",err.Error(),"; Entity:{{.table.Title}}")
                }
                return list
            }

            // Save {{.table.Comment}}
            func (t *{{.table.Title}}RepoImpl) Save(v *model.{{.table.Title}})(int,error){
                id,err := orm.Save(t._orm,v,int(v.{{title .table.Pk}}))
                if err != nil && err != sql.ErrNoRows{
                  log.Println("[ Orm][ Error]:",err.Error(),"; Entity:{{.table.Title}}")
                }
                return id,err
            }

            // Delete {{.table.Comment}}
            func (t *{{.table.Title}}RepoImpl) Delete(primary interface{}) error {
                err := t._orm.DeleteByPk(model.{{.table.Title}}{}, primary)
                if err != nil && err != sql.ErrNoRows{
                  log.Println("[ Orm][ Error]:",err.Error(),"; Entity:{{.table.Title}}")
                }
                return err
            }

            // Batch Delete {{.table.Comment}}
            func (t *{{.table.Title}}RepoImpl) BatchDelete(where string,v ...interface{})(int64,error) {
                r,err := t._orm.Delete(model.{{.table.Title}}{},where,v...)
                if err != nil && err != sql.ErrNoRows{
                  log.Println("[ Orm][ Error]:",err.Error(),"; Entity:{{.table.Title}}")
                }
                return r,err
            }

            `
)
