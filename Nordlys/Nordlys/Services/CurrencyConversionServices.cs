using Microsoft.EntityFrameworkCore;
using Nordlys.Data;

namespace Nordlys.Services;

/// <summary>
/// Currency conversion services implementation.
/// </summary>
public class CurrencyConversionServices : ICurrencyConversionServices
{
    private readonly ApplicationDbContext _applicationDb;

    public CurrencyConversionServices(ApplicationDbContext applicationDb)
    {
        _applicationDb = applicationDb;
    }
    
    public async Task<decimal> ConvertToEuro(Currency currency, decimal amount)
    {
        try
        {
            var latestReference = await _applicationDb.CurrencyConversions
                .Where(conversion => conversion.TargetId == currency.Id)
                .OrderByDescending(conversion => conversion.ConversionDate)
                .FirstAsync();

            return latestReference.EuroValue * amount;
        }
        catch (InvalidOperationException _)
        {
            return -1m;
        }
    }
}