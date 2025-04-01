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
    
}