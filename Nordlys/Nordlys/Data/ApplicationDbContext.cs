using Microsoft.EntityFrameworkCore;

namespace Nordlys.Data;

/// <summary>
/// Application Database context
/// </summary>
public class ApplicationDbContext : DbContext
{
    public ApplicationDbContext(DbContextOptions options) : base(options)
    {
    }
    
    public DbSet<Currency> Currencies { get; set; }
    public DbSet<CurrencyConversion> CurrencyConversions { get; set; }
}