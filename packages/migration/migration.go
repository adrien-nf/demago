package migration

type Migration interface {
	Up(builder *MigrationBuilder)
	Down(builder *MigrationBuilder)
}
