using Microsoft.EntityFrameworkCore.Migrations;

#nullable disable

namespace Nordlys.Data.Migrations
{
    /// <inheritdoc />
    public partial class Currency_symbols : Migration
    {
        /// <inheritdoc />
        protected override void Up(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.AddColumn<string>(
                name: "Symbol",
                table: "Currencies",
                type: "text",
                nullable: false,
                defaultValue: "");
        }

        /// <inheritdoc />
        protected override void Down(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DropColumn(
                name: "Symbol",
                table: "Currencies");
        }
    }
}
