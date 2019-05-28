/**
 * this file is auto generated by TTO-CodeGenerator v{{.global.Version}} !
 * if you want to modify this code,please read guide doc
 * and modify code template later.
 *
 * guide please see https://github.com/ixre/tto
 *
 */
!target:java/{{.global.Pkg}}/repo/{{.table.Title}}Repository.java
package {{pkg "java" .global.Pkg}}.repo;

import {{pkg "java" .global.Pkg}}.pojo.{{.table.Title}}Entity
import org.springframework.data.jpa.repository.JpaRepository
{{$pkType := type "java" .table.PkTypeId}}
/** {{.table.Comment}}仓储接口 */
public interface {{.table.Title}}Repository : JpaRepository<{{.table.Title}}Entity, {{$pkType}}>{

}
